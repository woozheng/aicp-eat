# aicp/llm.py
"""LLM — Multi-model chat interface with streaming support."""

import asyncio
import json
from typing import Dict, List, Optional, AsyncIterator, Union
from openai import AsyncOpenAI, AsyncStream


class LLM:
    def __init__(self, config: dict):
        self.default_model = config.get("models", {}).get("default", "gpt-3.5-turbo")
        self.providers = config.get("models", {}).get("providers", {})
        self._clients = {}
        self._model_to_client = {}
        self._model_configs = {}
        self.high_quality_model = None
        self.max_retries = config.get("models", {}).get("max_retries", 3)
        self.request_timeout = config.get("models", {}).get("request_timeout", 60)
        self.stream_timeout = config.get("models", {}).get("stream_timeout", 300)
        self._semaphore = asyncio.Semaphore(config.get("models", {}).get("max_concurrent", 10))
        self._init_clients()

    def _init_clients(self):
        for name, cfg in self.providers.items():
            api_key = cfg.get("api_key", "")
            base_url = cfg.get("base_url", "https://api.openai.com/v1")
            if not api_key or api_key.startswith("${"):
                print(f"  ⚠️  Skip {name}: API key not configured")
                continue

            client = AsyncOpenAI(
                api_key=api_key,
                base_url=base_url,
                timeout=self.request_timeout,
                max_retries=0  # We handle retries ourselves
            )
            self._clients[name] = client

            for model_cfg in cfg.get("models", []):
                model_id = model_cfg.get("id")
                self._model_configs[model_id] = {
                    "client": name,
                    "max_tokens": model_cfg.get("max_tokens", 4096),
                    "temperature": model_cfg.get("temperature", 0.7),
                    "supports_streaming": model_cfg.get("supports_streaming", True),
                }
                self._model_to_client[model_id] = name
                if model_cfg.get("default"):
                    self.default_model = model_id
                if model_cfg.get("high_quality") and not self.high_quality_model:
                    self.high_quality_model = model_id
                print(f"  ✅ {model_id} ({name})")

        if not self.default_model and self._model_to_client:
            self.default_model = next(iter(self._model_to_client))

        print(f"🤖 Default: {self.default_model} | Models: {len(self._model_to_client)}")

    def _get_client(self, model: str = None):
        model = model or self.default_model
        client_name = self._model_to_client.get(model)
        if client_name:
            return self._clients[client_name], model
        # Fallback to first available client
        if self._clients:
            name = next(iter(self._clients))
            return self._clients[name], model
        return None, model

    # ========================================================================
    # Non-streaming chat
    # ========================================================================
    async def _chat_impl(self, messages: List[Dict], model: str = None, **kwargs) -> str:
        model = model or self.default_model
        client, actual_model = self._get_client(model)
        if not client:
            return "[LLM not configured]"

        config = self._model_configs.get(actual_model, {})
        max_tokens = kwargs.pop("max_tokens", config.get("max_tokens", 4096))
        temperature = kwargs.pop("temperature", config.get("temperature", 0.7))

        for attempt in range(self.max_retries):
            try:
                print(f"🤖 LLM call (attempt {attempt + 1}/{self.max_retries}) → {actual_model}")

                resp = await asyncio.wait_for(
                    client.chat.completions.create(
                        model=actual_model,
                        messages=messages,
                        max_tokens=max_tokens,
                        temperature=temperature,
                        stream=False,
                        **kwargs
                    ),
                    timeout=self.request_timeout
                )

                content = resp.choices[0].message.content
                print(f"✅ Response: {len(content)} chars")
                return content

            except asyncio.TimeoutError:
                print(f"⏱️ Timeout (attempt {attempt + 1})")
                if attempt == self.max_retries - 1:
                    return "[LLM timeout]"
                await asyncio.sleep(min(2 ** attempt, 30))

            except Exception as e:
                error_msg = str(e)
                print(f"❌ LLM error: {type(e).__name__}: {error_msg[:100]}")

                if "rate_limit" in error_msg.lower() or "429" in error_msg:
                    wait = min(10 * (2 ** attempt), 120)
                    print(f"🔄 Rate limited, waiting {wait}s...")
                    await asyncio.sleep(wait)
                    continue

                if "connection" in error_msg.lower() or "timeout" in error_msg.lower():
                    if attempt == self.max_retries - 1:
                        return f"[LLM connection error: {error_msg[:100]}]"
                    await asyncio.sleep(2 ** attempt)
                    continue

                if attempt == self.max_retries - 1:
                    return f"[LLM error: {error_msg[:200]}]"
                await asyncio.sleep(1)

        return "[LLM max retries exceeded]"

    async def chat(self, messages: List[Dict], model: str = None, **kwargs) -> str:
        """Non-streaming chat. Returns full response."""
        async with self._semaphore:
            return await self._chat_impl(messages, model, **kwargs)

    # ========================================================================
    # Streaming chat
    # ========================================================================
    async def _chat_stream_impl(
        self, messages: List[Dict], model: str = None, **kwargs
    ) -> AsyncIterator[str]:
        model = model or self.default_model
        client, actual_model = self._get_client(model)
        if not client:
            yield "[LLM not configured]"
            return

        config = self._model_configs.get(actual_model, {})
        if not config.get("supports_streaming", True):
            # Fallback to non-streaming
            result = await self._chat_impl(messages, model, **kwargs)
            yield result
            return

        max_tokens = kwargs.pop("max_tokens", config.get("max_tokens", 4096))
        temperature = kwargs.pop("temperature", config.get("temperature", 0.7))

        for attempt in range(self.max_retries):
            stream = None
            try:
                print(f"🤖 LLM stream (attempt {attempt + 1}/{self.max_retries}) → {actual_model}")

                stream = await client.chat.completions.create(
                    model=actual_model,
                    messages=messages,
                    max_tokens=max_tokens,
                    temperature=temperature,
                    stream=True,
                    **kwargs
                )

                collected = []
                async for chunk in stream:
                    if chunk.choices and chunk.choices[0].delta.content:
                        content = chunk.choices[0].delta.content
                        collected.append(content)
                        yield content

                full_response = "".join(collected)
                print(f"✅ Stream complete: {len(full_response)} chars")
                return  # Success

            except asyncio.TimeoutError:
                print(f"⏱️ Stream timeout (attempt {attempt + 1})")
                if attempt == self.max_retries - 1:
                    yield "[LLM stream timeout]"
                    return
                await asyncio.sleep(min(2 ** attempt, 30))

            except Exception as e:
                error_msg = str(e)
                print(f"❌ Stream error: {type(e).__name__}: {error_msg[:100]}")

                if "rate_limit" in error_msg.lower() or "429" in error_msg:
                    wait = min(10 * (2 ** attempt), 120)
                    print(f"🔄 Rate limited, waiting {wait}s...")
                    await asyncio.sleep(wait)
                    continue

                if attempt == self.max_retries - 1:
                    yield f"[LLM stream error: {error_msg[:200]}]"
                    return
                await asyncio.sleep(1)

            finally:
                # Ensure stream is properly closed
                if stream and hasattr(stream, 'close'):
                    try:
                        await stream.close()
                    except Exception:
                        pass

        yield "[LLM max retries exceeded]"

    async def chat_stream(
        self, messages: List[Dict], model: str = None, **kwargs
    ) -> AsyncIterator[str]:
        """Streaming chat. Yields tokens as they arrive."""
        async with self._semaphore:
            async for token in self._chat_stream_impl(messages, model, **kwargs):
                yield token

    # ========================================================================
    # JSON mode (non-streaming only)
    # ========================================================================
    async def chat_json(self, messages: List[Dict], model: str = None, **kwargs) -> dict:
        """Call LLM and parse response as JSON. Retries on parse failure."""
        for attempt in range(3):
            raw = await self.chat(messages, model, **kwargs)
            raw = raw.strip()

            # Extract JSON from markdown code blocks
            if "```json" in raw:
                raw = raw.split("```json")[1].split("```")[0].strip()
            elif "```" in raw:
                parts = raw.split("```")
                if len(parts) > 1:
                    raw = parts[1].split("```")[0].strip()

            try:
                return json.loads(raw)
            except json.JSONDecodeError as e:
                print(f"⚠️ JSON parse failed (attempt {attempt + 1}/3): {e}")
                if attempt == 2:
                    return {"content": raw, "parse_error": str(e)}

                # Retry with stricter prompt
                messages.append({
                    "role": "system",
                    "content": "Respond ONLY with valid JSON. No markdown, no explanation."
                })

        return {"content": raw}

    # ========================================================================
    # Health check
    # ========================================================================
    async def health_check(self) -> bool:
        """Check if at least one client is responsive."""
        for client in self._clients.values():
            try:
                await asyncio.wait_for(client.models.list(), timeout=5)
                return True
            except Exception:
                continue
        return False