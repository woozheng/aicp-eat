# AICP-EAT

## Eat Every Python Library → HTTP API.

```bash
pip install aicp-eat
```
## Eat whatever you want

```bash
pip install pandas
python eater/eat.py pandas            # Standard library — feed it to eat.py

pip install numpy
python eater/eat.py numpy             # Standard library — feed it to eat.py

pip install pillow
python eater/eat_pil.py               # Nested modules — throw it to AI, let it generate a custom script

pip install langchain-community
python eater/eat_langchain.py         # Nested modules — throw it to AI, let it generate a custom script
```

```text
plugins/
├── pandas.py              → 102 functions
├── numpy.py               → 460 functions
├── PIL.py                 → 664 functions
└── langchain_community.py → 7512 functions   ← new number every day. chaos.
────────────────────────────────────────
Total: 8738 functions. 4 files.
```
and more. Come eat.

## Run

```bash
# Set your API key in aicp.yaml, or:
export AGGREGATOR_API_KEY="sk-your-key"  # Windows: setx AGGREGATOR_API_KEY "sk-your-key"

aicp
```
```bash
curl -X POST http://localhost:9000/api/pandas \
  -H "Content-Type: application/json" \
  -d '{"function":"array","args":{"data":[1,2,3,4,5]}}'
```
Return

```json
{
  "result": "IntegerArray\n[1, 2, 3, 4, 5]\nLength: 5, dtype: Int64",
  "source": "pandas"
}
```
Open http://localhost:9000, browse the API menu at /api/pandas, /api/numpy, /api/PIL, /api/langchain_community. Each one has a help() with all functions listed.


## [Why eat the entire Python ecosystem? →](/docs/Journey.md)

Go calls pandas. Rust calls numpy. Frontend calls LangChain. Any language, HTTP everything.

## Eat Everything to HTTP

| Language | Status | Engine |
|----------|--------|--------|
| Python | ✅ Eaten | [aicp-eat](https://github.com/woozheng/aicp-eat) |
| Go | ✅ Eaten | [go/](./go/) |
| Rust | ✅ Eaten | [rust/](./rust/) |
| JavaScript | 🪑 Feed AI | Drop the protocol. AI writes it. |
| Java | 🪑 Feed AI | Drop the protocol. AI writes it. |
| Zig | 🪑 Feed AI | Drop the protocol. AI writes it. |
| ... | 🪑 Feed AI | You get it. |

**Same protocol. Same call.**

```bash
# Python
curl localhost:9000/api/numpy -d '{"meta":{"function":"std"},"payload":{"args":{"a":[1,2,3,4,5]}}}'

# Go
curl localhost:9000/api/github_com_gin-gonic_gin -d '{"meta":{"function":"Mode"},"payload":{"args":{}}}'

# Rust
curl localhost:9000/api/serde -d '{"meta":{"function":"main"},"payload":{"args":{}}}'


```
AI reads the protocol. AI writes the engine. AI writes the eater. AI generates everything.

80 lines. 3 atoms. Any language.

## AI-Generated Experiments 


AI read the AICP protocol and generated complete systems across seven domains. Each from a single human sentence.

> **Want to generate systems like these? → [Start here](https://github.com/woozheng/aicp)**

| Experiment | Domain | Human Said |
|---|---|---|
| [aicp-os-kernel](https://github.com/woozheng/aicp-os-kernel) | OS Kernel | Microkernel OS |
| [aicp-quantum](https://github.com/woozheng/aicp-quantum) | Quantum Computing | Quantum computing simulator |
| [aicp-protein](https://github.com/woozheng/aicp-protein) | Protein Folding | Protein folding |
| [aicp-llm-trainer](https://github.com/woozheng/aicp-llm-trainer) | LLM Training | LLM distributed training |
| [aicp-riemann](https://github.com/woozheng/aicp-riemann) | Riemann Hypothesis | How to approach and verify the Riemann Hypothesis |
| [aicp-ai-chip](https://github.com/woozheng/aicp-ai-chip) | AI Chip | AI intelligent chip computing system |
| [aicp-raw-experiments](https://github.com/woozheng/aicp-raw-experiments) | Raw Experiments | More raw experiments |

[→ More about the protocol](/docs/AICP_Protocol_v3.md)

## Dependencies

- Python >= 3.10
- aiohttp >= 3.9
- openai >= 1.0
- pyyaml >= 6.0

*Eat it, Eat it, Eat it, Eat it*  
*No one wants to be defeated*  
*Showin' how funky and strong is your fight*  
*It doesn't matter who's wrong or right*  
*Just eat it, eat it*  
*Just eat it, eat it*

---
[See More →](/docs/Journey.md)

## [MIT License](LICENSE) · Dvwoo