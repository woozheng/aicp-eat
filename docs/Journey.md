## My Journey

I first devoured LangChain — 999 APIs. Spectacular.

Then I tested one: WikipediaLoader. It just calls the `wikipedia` library underneath.

Another: ChatOpenAI. It just calls the `openai` library.

I stopped. Why did I eat 999 wrapper classes?

I ate the bottom instead:

```bash
python eater/eat.py requests
python eater/eat.py wikipedia
python eater/eat.py openai
```

AI reads their help(). Orchestrates. 3 APIs become infinite combinations.

LangChain took 999 classes to do the same thing.
Bottom libraries need 3 HTTP APIs.
AI orchestrates. Done.

The 999 plugins? Not deleted. In `stores/`.

LangChain made Python libraries easier to use. That's real value.
It just turned out: when AI orchestrates, going direct is even easier.

So I kept them. As a reminder.
Eat the wrappers. Eat the bottom. Both work.
See what suits you.

Why I Didn't Eat AutoGen.

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


