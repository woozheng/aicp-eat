# plugins/dev.py
"""
dev — AI App Generator + Orchestrator
Route: POST /api/dev
"""

import json
import uuid
from pathlib import Path
import importlib.util

import core

PLUGINS_DIR = Path("plugins")
WWW_DIR = Path("www")
GUIDE_FILE = Path("docs/AICP_plugin_guide.md")
HISTORY_FILE = Path("data") / "dev_history.json"

_guide_cache = None

# Built-in fallback guide placeholder
# Copy your full guide text here if docs/ file not found
FALLBACK_GUIDE = """
# AICP Plugin Guide v2.1
## Rule
```python

async def execute(envelop, agent):
    data = envelop.payload.get("content", "")
    llm = agent.llm
    result = await llm.chat([...])
    envelop.payload["result"] = result
    return envelop
```    
Input: envelop.payload → Process: LLM or logic → Output: envelop.payload

### Route = File Path
```text
plugins/hello.py           → POST /api/hello
plugins/recruit/parse.py   → POST /api/recruit/parse
plugins/a/b/c.py           → POST /api/a/b/c
Copy file = route registered. No config needed.
```
### Plugin Help (Recommended)

Every plugin should include a `help()` function so other plugins and AI know how to call it.
## Templates



### Standard (translate, Q&A, classify, chat)
```python
def help():
    return {
        "route": "/api/translate",
        "input": {"content": "text to translate"},
        "output": {"result": "translated text"},
        "description": "Translate text to English"
    }
async def execute(envelop, agent):
    try:
        llm = agent.llm
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        result = await llm.chat([
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": envelop.payload.get("content", "")}
        ])

        envelop.payload["result"] = result.strip()
    except Exception as e:
        envelop.payload = {"error": str(e)}
    return envelop
```    
### Pipeline (analyze → process → summarize)
```python

def help():
    return {
        "route": "/api/xxx",
        "input": {"content": "text"},
        "output": {"keywords": [...], "summary": {...}},
        "description": "Multi-stage analysis pipeline"
    }
async def execute(envelop, agent):
    try:
        llm = agent.llm
        text = envelop.payload.get("content", "")

        # Step 1: Extract keywords
        keywords = await llm.chat_json([
            {"role": "system", "content": "Extract 5 keywords. Return JSON: {keywords: [...]}"},
            {"role": "user", "content": text}
        ])

        # Step 2: Search based on keywords
        search_result = await llm.chat([
            {"role": "system", "content": f"Based on keywords {keywords}, find relevant information."},
            {"role": "user", "content": text}
        ])

        # Step 3: Summarize findings
        summary = await llm.chat_json([
            {"role": "system", "content": "Summarize findings. Return JSON: {summary, confidence}"},
            {"role": "user", "content": search_result}
        ])

        envelop.payload = {"keywords": keywords, "summary": summary}
    except Exception as e:
        envelop.payload = {"error": str(e)}
    return envelop
```
### Parallel (multi-expert, batch)
```python
import asyncio

def help():
    return {
        "route": "/api/xxx",
        "input": {"content": "text"},
        "output": {"results": ["opinion1", "opinion2", ...]},
        "description": "Multi-expert parallel analysis"
    }
async def execute(envelop, agent):
    try:
        llm = agent.llm
        task = envelop.payload.get("content", "")

        async def expert(i):
            return await llm.chat([
                {"role": "system", "content": f"You are expert {i+1}. Give a unique opinion."},
                {"role": "user", "content": task}
            ])

        results = await asyncio.gather(*[expert(i) for i in range(3)])

        envelop.payload = {"results": results}
    except Exception as e:
        envelop.payload = {"error": str(e)}
    return envelop
```    
### No LLM (webhook, proxy, data processing)
```python
def help():
    return {
        "route": "/api/xxx",
        "input": {"content": "data"},
        "output": {"received": "data", "status": "ok"},
        "description": "Data processor"
    }
async def execute(envelop, agent):
    try:
        data = envelop.payload.get("content", envelop.payload)
        result = {"received": data, "status": "ok"}
        envelop.payload = result
    except Exception as e:
        envelop.payload = {"error": str(e)}
    return envelop
```    
## Quick Match
| Keywords | Template |
|----------|----------|
| translate, Q&A, classify, chat, summarize | Standard |
| analyze → process → summarize, step by step | Pipeline |
| parallel, multi-expert, batch, simultaneous | Parallel |
| webhook, proxy, forward, no LLM | No LLM |

## Output Format
```text
=== PLUGIN: plugins/{endpoint}.py ===
... code ...
=== HTML: www/{page}.html ===
... html ...
=== END ===
No HTML? Write === HTML: null ===.
```
## Memory (optional)
```python
import json
from pathlib import Path

async def execute(envelop, agent):
    memory_file = Path("data") / "memory.json"
    history = json.loads(memory_file.read_text()) if memory_file.exists() else []
    
    # ... process ...
    
    history.append({"input": text, "output": result})
    memory_file.write_text(json.dumps(history[-20:]))
```    
## Checklist
- [ ] `execute(envelop, agent)` only
- [ ] `try-except`
- [ ] Check `agent.llm` before use
- [ ] Return envelop
- [ ] File path = route

Read. Pick template. Generate code. Done.
"""


def help():
    return {
        "route": "/api/dev",
        "input": {"content": "app description"},
        "output": {"ok": True, "plugin_route": "/api/xxx", "page_url": "xxx.html"},
        "description": "Generate AICP app + frontend from natural language. Can orchestrate existing APIs."
    }


def load_guide():
    global _guide_cache
    if _guide_cache is None:
        if GUIDE_FILE.exists():
            _guide_cache = GUIDE_FILE.read_text(encoding="utf-8")
        else:
            _guide_cache = FALLBACK_GUIDE
    return _guide_cache


def load_history():
    if HISTORY_FILE.exists():
        return json.loads(HISTORY_FILE.read_text(encoding="utf-8"))
    return []


def save_history(item):
    HISTORY_FILE.parent.mkdir(parents=True, exist_ok=True)
    hist = load_history()
    hist.append(item)
    if len(hist) > 50:
        hist = hist[-50:]
    HISTORY_FILE.write_text(json.dumps(hist, ensure_ascii=False, indent=2), encoding="utf-8")


def collect_api_docs():
    """Scan all plugins and collect help() info."""
    docs = []
    for name in core.plugins:
        if name == "dev":
            continue
        try:
            mod_name = f"plugins.{name.replace('/', '.')}"
            mod = __import__(mod_name, fromlist=["help"])
            if hasattr(mod, "help"):
                docs.append(mod.help())
        except Exception:
            pass
    return docs


async def execute(envelop, agent):
    payload = envelop.payload
    user_input = payload.get("content", "").strip()

    if not user_input:
        envelop.payload = {"ok": False, "error": "Please describe the app you want"}
        return envelop

    guide = load_guide()
    if not guide:
        envelop.payload = {"ok": False, "error": "Guide not available"}
        return envelop

    llm = agent.llm if agent else None
    if not llm:
        envelop.payload = {"ok": False, "error": "LLM not available"}
        return envelop

    # Collect existing APIs for orchestration
    api_docs = collect_api_docs()
    api_info = ""
    if api_docs:
        api_info = "\n\n## Available APIs You Can Orchestrate\n"
        for api in api_docs[:50]:
            api_info += f"- {api.get('route')}: {api.get('description', '')}\n"
        api_info += "\nTo call these APIs from your plugin, use aiohttp:\n"
        api_info += "import aiohttp\n"
        api_info += "async with aiohttp.ClientSession() as session:\n"
        api_info += "    async with session.post('http://localhost:9000/api/xxx', json={'content': '...'}) as resp:\n"
        api_info += "        data = await resp.json()\n"

    system_prompt = f"""You are an AICP app generator. Create complete, runnable code.

Follow these rules:
{guide}
{api_info}

【Output Format】
=== PLUGIN: plugins/xxx.py ===
Complete code (no markdown blocks)
=== HTML: www/xxx.html ===
Complete HTML (no markdown blocks)
=== END ===

【Critical】
1. MUST end with === END ===
2. No HTML? Write === HTML: null ===
3. Code directly, no ```python or ```html wrappers
4. Every plugin MUST have help() function
5. For multi-step tasks, orchestrate existing APIs using aiohttp
6. HTML filename: use simple name like www/xxx.html, NOT subdirectories
7. Complete output, one shot"""

    try:
        raw = await llm.chat([
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": user_input}
        ])
    except Exception as e:
        envelop.payload = {"ok": False, "error": f"LLM call failed: {e}"}
        return envelop

    # Parse
    result = {"plugin_file": None, "plugin_code": None, "html_file": None, "html_code": None}
    lines = raw.split("\n")
    mode = None
    current_file = None
    current_code = []

    for line in lines:
        stripped = line.strip()

        if stripped in ("```python", "```html", "```", "```json"):
            continue

        if stripped.startswith("=== PLUGIN:") and "===" in stripped:
            _save_block(mode, current_file, current_code, result)
            mode = "plugin"
            current_file = stripped.replace("=== PLUGIN:", "").replace("===", "").strip()
            current_code = []

        elif stripped.startswith("=== HTML:") and "===" in stripped:
            _save_block(mode, current_file, current_code, result)
            mode = "html"
            current_file = stripped.replace("=== HTML:", "").replace("===", "").strip()
            current_code = []

        elif stripped == "=== END ===":
            _save_block(mode, current_file, current_code, result)
            mode = None
            current_file = None
            current_code = []

        elif mode and current_file:
            current_code.append(line)

    _save_block(mode, current_file, current_code, result)

    # Write plugin
    plugin_file = result["plugin_file"]
    plugin_code = result["plugin_code"]
    plugin_route = None

    if plugin_file and plugin_code:
        filepath = Path(plugin_file)
        filepath.parent.mkdir(parents=True, exist_ok=True)
        filepath.write_text(plugin_code, encoding="utf-8")
        try:
            route_path = str(filepath.relative_to(PLUGINS_DIR)).replace("\\", "/").replace(".py", "")
        except ValueError:
            route_path = filepath.stem
        plugin_route = f"/api/{route_path}"

        # Hot-load
        spec = importlib.util.spec_from_file_location(f"plugin_{route_path.replace('/', '_')}", filepath)
        mod = importlib.util.module_from_spec(spec)
        try:
            spec.loader.exec_module(mod)
            if hasattr(mod, "execute"):
                core.plugins[route_path] = mod.execute
                print(f"✅ App: {plugin_route}")
        except Exception as e:
            print(f"⚠️ App saved but import failed: {e}")

    # Write HTML
    html_file = result["html_file"]
    html_code = result["html_code"]
    page_url = None

    if html_file and html_code:
        filepath = Path(html_file)
        filepath.parent.mkdir(parents=True, exist_ok=True)
        filepath.write_text(html_code, encoding="utf-8")
        page_url = f"/{filepath.name}"
        print(f"✅ Frontend: {page_url}")

    save_history({
        "id": uuid.uuid4().hex[:8],
        "prompt": user_input,
        "plugin_route": plugin_route,
        "page_url": page_url,
    })

    envelop.payload = {
        "ok": True,
        "plugin_route": plugin_route,
        "page_url": page_url,
    }
    return envelop


def _save_block(mode, current_file, current_code, result):
    if not mode or not current_file or not current_code:
        return
    code = "\n".join(current_code).strip()
    if not code:
        return
    if mode == "plugin":
        result["plugin_file"] = current_file
        result["plugin_code"] = code
    elif mode == "html" and current_file != "null":
        result["html_file"] = current_file
        result["html_code"] = code