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