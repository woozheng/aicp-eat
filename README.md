# AICP-EAT

## Eat Every Python Library → HTTP API.

```bash
pip install aicp-eat
```
## Eat whatever you want

```bash
pip install pandas
python eater/eat.py pandas        # → 59 APIs

pip install numpy
python eater/eat.py numpy          # → 92 APIs

pip install langchain-community
python eater/eat_langchain.py      # → 999 APIs
```
## Run

```bash
# Set your API key in aicp.yaml, or:
export AGGREGATOR_API_KEY="sk-your-key"  # Windows: setx AGGREGATOR_API_KEY "sk-your-key"

aicp
```
```bash
curl http://localhost:9000/api/pandas/array -d '{"params":{"data":[1,2,3,4,5]}}'
```
Return

```json
{
  "result": "IntegerArray\n[1, 2, 3, 4, 5]\nLength: 5, dtype: Int64",
  "source": "pandas"
}
```
Open http://localhost:9000, test all APIs in the workbench.


## [Why eat the entire Python ecosystem? →](/docs/Journey.md)

Go calls pandas. Rust calls numpy. Frontend calls LangChain. Any language, HTTP everything.

```text
pandas    → 59 APIs
numpy     → 92 APIs
LangChain → 999 APIs
sklearn   → 300+ APIs
scipy     → 107 APIs
PIL       → 471 APIs
requests  → 9 APIs
─────────────────────
2000+ APIs. HTTP.
```
and more. Come eat.

[Plugin Store →](stores/README.md)

---
## AI reads help(). AI orchestrates. 80 lines engine.AI-Generated Experiments 

AI read the AICP protocol and generated complete systems across seven domains. Each from a single human sentence.

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