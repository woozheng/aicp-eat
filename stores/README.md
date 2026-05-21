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


But honestly? Don't just eat what we ate. Eat what YOU want.

You're a Python developer. You know exactly which 10 functions you actually use.
The 2000+ APIs are just to show you it's possible. Your real treasure is the handful of libraries you import every single day.

### Here's the real move:
Take eater/eat.py, throw it to any AI, and say:
```text
"Use this template. Generate a plugin for just these three functions: scipy.optimize.minimize, sklearn.ensemble.RandomForestClassifier, and PIL.Image.filter. Don't eat the whole library. Just these."
```
The AI reads the script, understands the pattern, and generates exactly what you need. You don't need to understand how eat.py works. You just need to know what you want to eat.

### That's the power.
We gave you 2000+ APIs to prove it works. But the real move is you, telling AI what to eat, and AI doing it in seconds.

Go eat your own favorites. That's where the magic is.

## Polish descriptions

```bash
python optimize_descriptions.py sklearn
```
Let AI read help() and rewrite descriptions in plugins/sklearn/ — so AI orchestrators understand exactly what each API does and how to call it. Or use your own optimization method. Edit LLM_CONFIG in the script first and customize the prompt to your liking.