# AICP

## Eat Every Python Library → HTTP API.

```bash
pip install aicp-eat
# Set your API key in aicp.yaml, or:
export AGGREGATOR_API_KEY="sk-your-key"  # Windows: setx AGGREGATOR_API_KEY "sk-your-key"
aicp
```
## Open http://localhost:9000, test all APIs in the workbench.

```bash
# Eat pandas → 59 APIs
python eater/eat.py pandas
curl http://localhost:9000/api/pandas/read_csv -d '{"params":{"filepath":"data.csv"}}'

# Eat numpy → 100+ APIs  
python eater/eat.py numpy

# Eat LangChain → 999 APIs
python eater/eat_langchain.py
```

## Go calls pandas. Rust calls numpy. Frontend calls LangChain. Any language, HTTP everything.

```text
pandas    → 59 APIs
numpy     → 100+ APIs
LangChain → 999 APIs
sklearn   → 301 APIs
scipy     → 107 APIs
PIL       → 471 APIs
requests  → 9 APIs
─────────────────────
2000+ APIs. HTTP.
```
## AI reads help(). AI orchestrates. 80 lines engine.
---
## Dependencies

- Python >= 3.10
- aiohttp >= 3.9
- openai >= 1.0
- pyyaml >= 6.0

### *Just eat it, eat it, eat it, eat it*
### *No one wants to be defeated*
### *Showin' how funky and strong is your fight*
### *It doesn't matter who's wrong or right*
### *Let's eat it, eat it*
### *Let's eat it, eat it*

---
[See More →](/docs/Journey.md)

## [MIT License](LICENSE) · Dvwoo