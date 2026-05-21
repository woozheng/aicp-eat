AICP Plugin Development Guide v2.1
Overview
1 standard template (80% of use cases) + 3 pattern variants (15%) = 95% coverage.

Key rule: Plugins contain only the execute function. HTTP routes auto-registered by Gateway.

Standard Template (80% of use cases)
Use when: Translate, summarize, sentiment, Q&A, classify, code generation — receive input → call LLM once → return output.

python
"""
{PluginName} — {Description}
Route: POST /api/{endpoint}
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        user_input = payload.get("content", "")

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        result = await llm.chat([
            {"role": "system", "content": "{system_prompt}"},
            {"role": "user", "content": user_input}
        ])

        envelop.payload = {"result": result.strip()}
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop
Placeholders
Placeholder	Meaning	Example
{PluginName}	Plugin name	Sentiment Analyzer
{Description}	One-line description	Analyze sentiment of user input
{endpoint}	API route path	sentiment
{system_prompt}	LLM role prompt	You are a sentiment analysis expert
"content"	Input field name	"text" / "message" / "query"
"result"	Output field name	"analysis" / "output" / "response"
Filename: plugins/{endpoint}.py
Test:
bash
curl -X POST http://127.0.0.1:9000/api/{endpoint} \
  -H "Content-Type: application/json" \
  -d '{"content":"test"}'
Pattern 1: Multi-Stage Pipeline
Use when: Evaluate → Process → Summarize. Multiple sequential LLM calls.

python
"""
{PluginName} — Multi-stage Pipeline
Route: POST /api/{endpoint}
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        user_input = payload.get("content", "")

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        # Phase 1: Analyze
        analysis = await llm.chat_json([
            {"role": "system", "content": "Analyze the input. Return JSON: {key_points, sentiment}"},
            {"role": "user", "content": user_input}
        ])

        # Phase 2: Respond
        response = await llm.chat([
            {"role": "system", "content": "Generate a response based on the analysis."},
            {"role": "user", "content": f"Analysis: {analysis}\nInput: {user_input}"}
        ])

        # Phase 3: Summarize
        summary = await llm.chat_json([
            {"role": "system", "content": "Summarize. Return JSON: {summary, action_items}"},
            {"role": "user", "content": response}
        ])

        envelop.payload = {
            "analysis": analysis,
            "response": response.strip(),
            "summary": summary
        }
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop
Method	Returns
llm.chat()	str
llm.chat_json()	dict
Pattern 2: Parallel Processing
Use when: Multiple expert opinions, batch analysis, simultaneous LLM calls.

python
"""
{PluginName} — Parallel Processing
Route: POST /api/{endpoint}
"""

import asyncio

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        task = payload.get("content", "")
        worker_count = payload.get("workers", 3)

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        async def worker(i):
            result = await llm.chat([
                {"role": "system", "content": f"You are expert {i+1}. Give a unique opinion."},
                {"role": "user", "content": task}
            ])
            return {"expert": i+1, "opinion": result.strip()}

        results = await asyncio.gather(*[worker(i) for i in range(worker_count)])

        envelop.payload = {"results": results}
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop
Pattern 3: Pure Data Processing (No LLM)
Use when: Webhook receiver, log collector, data validation, request proxy. No LLM needed.

python
"""
{PluginName} — Data Processor (No LLM)
Route: POST /api/{endpoint}
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        data = payload.get("content", payload)

        # Your business logic here
        result = {"received": data, "status": "processed"}

        envelop.payload = result
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop
Quick Decision Table
User keywords	Use pattern
translate, summarize, sentiment, Q&A, classify	Standard
evaluate, analyze, multi-stage, step by step, pipeline	Pattern 1
parallel, simultaneous, multi-expert, batch	Pattern 2
forward, webhook, collect, no LLM, proxy	Pattern 3
Plugin Rules
1. Single Entry Point
python
async def execute(envelop, agent):
Two parameters. Always.

2. Access Resources Through agent
Need	Use
LLM	agent.llm
Message bus	agent.bus
Agent ID	agent.id
agent supports both agent.llm and agent.get("llm").

3. Modify payload, Return envelop
python
envelop.payload["result"] = "done"
return envelop
4. Route = File Path
File path	API route
plugins/hello.py	POST /api/hello
plugins/recruit/parse.py	POST /api/recruit/parse
plugins/ticket/classify.py	POST /api/ticket/classify
Subdirectories become path segments. No manual registration.

5. Plugin Returns Plain Data
python
# Plugin returns:
envelop.payload = {"result": "Hello"}

# Client receives:
{"result": "Hello"}
6. Python 3.10+ Compatible
No backslashes in f-strings. No match/case. No | in except.

Checklist
Only execute function

try-except error handling

LLM availability check (if using LLM)

File saved to plugins/{endpoint}.py

Route matches file path

AI reads this guide, selects pattern by keywords, fills placeholders, generates correct AICP plugin.