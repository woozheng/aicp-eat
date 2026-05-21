# Plugin Store

Pre-eaten plugin packs. Download, unzip, drop into `plugins/`. Done.

## Usage

```bash
pip install pandas  # or whatever you need
# Download the pack → unzip to plugins/ → restart AICP
```
## Packs

| Pack | APIs | Eat |
|------|------|-----|
| [999_langchain_plugins.zip](./python-3.11.7/999_langchain.zip) | 999 | ✅ Python 3.11.7 |
| [59_pandas_plugins.zip](./python-3.11.7/59_pandas.zip) | 59 | ✅ Python 3.11.7 |
| [92_numpy_plugins.zip](./python-3.11.7/92_numpy.zip) | 92 | ✅ Python 3.11.7 |
| [301_sklearn_plugins.zip](./python-3.11.7/302_sklearn.zip) | 302 | ✅ Python 3.11.7 |
| [107_scipy_plugins.zip](./python-3.11.7/107_scipy.zip) | 107 | ✅ Python 3.11.7 |
| [471_pil_plugins.zip](./python-3.11.7/471_pil.zip) | 471 | ✅ Python 3.11.7 |
| [9_requests_plugins.zip](./python-3.11.7/9_requests.zip) | 9 | ✅ Python 3.11.7 |
**2000+ APIs. Eat free. Run depends on your own env.**


## Eat your own
```bash
python eater/eat.py pandas        # eat the whole library
python eater/eat.py mylib /path   # eat from source
```
## Want specific functions only? Throw eater/eat.py to AI:

"Use this template, generate a script for just pandas read_csv and to_csv"

## Polish descriptions

```bash
python optimize_descriptions.py sklearn
```
Let AI read help() and rewrite descriptions in plugins/sklearn/ — so AI orchestrators understand exactly what each API does and how to call it. Or use your own optimization method. Edit LLM_CONFIG in the script first and customize the prompt to your liking.