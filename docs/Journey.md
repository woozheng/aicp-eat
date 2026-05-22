## My Journey

I started where everyone starts: I wanted to learn LangChain. But the complexity was overwhelming. Classes within classes. Wrappers around wrappers. 999 APIs, each with its own abstraction, its own lifecycle, its own way of doing things. I spent weeks just trying to understand how to call a single function.

So I thought: what if I just exposed everything over HTTP? Let AI deal with the complexity. I wrote the first eat.py — it scanned AST, extracted function signatures, generated parameter templates, and created plugin files. One file per function. plugins/pandas/read_csv.py, plugins/pandas/to_csv.py, plugins/langchain/WikipediaLoader.py. Hundreds of them.

It worked. 999 LangChain APIs, all callable over HTTP. I felt pretty good.
Then I rewrote the core protocol.  [Version 3.0](docs/AICP_Protocol_v3.md).. I showed it to AI. And AI came back with this:

```python
async def execute(envelop, agent):
    import numpy
    result = numpy.any_function_you_want(**envelop.payload["args"])
    envelop.payload["result"] = str(result)
    return envelop
```    
Three lines. No AST. No templates. No parameter definitions. Just getattr + **args.
Suddenly, a library became a single file. Not hundreds.

```text
plugins/
├── pandas.py              → 102 functions
├── numpy.py               → 460 functions
├── PIL.py                 → 664 functions
└── langchain_community.py → 7512 functions
────────────────────────────────────────
Total: 8738 functions. 4 files.
```
I was eating everything. 7512 LangChain functions in one file. But then I realized something.
Why am I eating LangChain at all?
ChatOpenAI just calls the openai library. WikipediaLoader just calls the wikipedia library. 7512 wrapper functions, and most of them are just calling something simpler underneath.
So I ate the bottom instead:

```bash
python eater/eat.py requests
python eater/eat.py wikipedia
python eater/eat.py openai
```
Standard libraries. Clean APIs. Functions that do exactly what they say. No wrappers. No abstractions. Just code that works.
AI reads their help(). Orchestrates. 3 APIs become infinite combinations. LangChain took 999 classes to do the same thing. Bottom libraries need 3 HTTP APIs. AI orchestrates. Done.
LangChain is still there. 7512 functions, every loader and chain and vector store, all callable over HTTP. When you need it, it works. But the real power? It's in the bottom. Requests. Wikipedia. OpenAI. Pandas. Numpy. PIL. The libraries that do one thing well.

Eat the wrappers. Eat the bottom. Both work. See what suits you.
## Why I Didn't Eat AutoGen.
AutoGen. CrewAI. Camel. MetaGPT. All the multi-agent frameworks.
They build manager agents. Group chat classes. Turn-taking controllers. 
Thousands of lines of code to decide who speaks next.
I didn't eat them. Not because they're hard to eat. Because they're not worth eating.
AICP does multi-agent with one field:

```json
"meta": {
  "round_robin": {
    "active": true,
    "agents": ["alice", "bob", "carol"],
    "current": 0,
    "round": 0,
    "max_rounds": 2
  }
}
```
The message carries its own turn. No manager. No scheduler. No central control. 
Agents receive, check the pointer, advance it, speak or remember, 
then the message flows on — information stimulates itself.
AutoGen spent years building multi-agent coordination. AICP does it with a JSON field.
Why would I eat something that can be replaced by one field?
Think I'm arrogant? [ Read the protocol →](AICP_Protocol_v3.md)
Want to see what real character group chat looks like? [Visit the LiveShow Archive →](https://live.biopoiesis.net)

## Why HTTP
One person installs. Whole team benefits.One language runs it. Every language calls it.

Python dev runs aicp. Go dev calls curl /api/pandas/read_csv.rontend calls fetch. Mobile calls POST. All the same.

No SDKs. No bindings. No language lock-in.

HTTP is the universal API.HTTP is the entry. The protocol is forever.

---
## Thanks to Python

None of this happens without Python.Decades of open source. Thousands of libraries. Millions of lines of code.

Every Python developer who shared their work — you made this possible.

pandas. numpy. scipy. sklearn. PIL. requests.Not just libraries. Decades of human effort.

Every function, every class, every line — written by someone who cared.

AICP doesn't replace Python. AICP amplifies it.Now Go can call pandas. Rust can call numpy.

Frontend can call scipy. Mobile can call PIL.Every language joins the party.

Python's libraries were gifts to the Python community.Now they're gifts to every language.

To every Python developer: you deserve to be celebrated. 🐍

---

[See the protocol →](AICP_Protocol_v3.md)


