# eater/eat_langchain.py
"""
LangChain Eater — Devour langchain_community into AICP plugins
Usage: python eater/eat_langchain.py
"""

import ast
from pathlib import Path

OUTPUT_DIR = Path("plugins") / "langchain"

METHOD_MAP = {
    "Loader": "load", "Reader": "load_data",
    "Tool": "run", "Wrapper": "run",
    "Chat": "invoke", "LLM": "invoke",
    "VectorStore": "similarity_search",
    "Retriever": "get_relevant_documents",
    "Embeddings": "embed_query",
    "Classifier": "predict", "Regressor": "predict",
    "KMeans": "fit_predict", "Scaler": "transform",
}

SKIP_PATTERNS = [
    "Error", "Exception", "Warning", "Base", "Meta", "Abstract",
    "Mixin", "Model", "Type", "Ops", "Impl", "Helper", "Util",
    "Internal", "Protocol", "Interface", "Test",
]

PLUGIN_TPL = '''"""
{description}
Route: POST /api/langchain/{name}
"""

import inspect
from {module_path} import {class_name}

def help():
    try:
        sig = inspect.signature({class_name}.__init__)
        params = {{}}
        for name, param in sig.parameters.items():
            if name == "self":
                continue
            if param.default is inspect.Parameter.empty:
                params[name] = "<required>"
            else:
                params[name] = str(param.default)
    except:
        params = {{"content": "input"}}
    
    return {{
        "route": "/api/langchain/{name}",
        "input": {{"params": params}},
        "output": {{"result": "output"}},
        "description": "{description}"
    }}

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        content = payload.get("content", "")
        params = payload.get("params", {{}})
        
        try:
            if params:
                instance = {class_name}(**params)
            else:
                instance = {class_name}({ctor_args})
        except:
            try:
                instance = {class_name}()
            except:
                instance = {class_name}(content)
        
        result = None
        if hasattr(instance, "{method}"):
            result = instance.{method}({call_args})
        elif hasattr(instance, "run"):
            result = instance.run({call_args})
        elif hasattr(instance, "invoke"):
            result = instance.invoke({call_args})
        elif hasattr(instance, "load"):
            result = instance.load()
        elif hasattr(instance, "predict"):
            result = instance.predict(content)
        elif hasattr(instance, "transform"):
            result = instance.transform(content)
        elif hasattr(instance, "fit_predict"):
            result = instance.fit_predict(content)
        elif hasattr(instance, "similarity_search"):
            result = instance.similarity_search(content)
        elif hasattr(instance, "load_data"):
            result = instance.load_data()
        else:
            result = str(instance)
        
        if hasattr(result, "content"):
            result = result.content
        elif hasattr(result, "page_content"):
            result = result.page_content
        elif isinstance(result, list):
            result = [str(r) for r in result[:10]]
        elif isinstance(result, dict):
            result = result.get("output", result.get("result", str(result)))
        else:
            result = str(result)
        
        envelop.payload["result"] = result
        envelop.payload["source"] = "langchain"
    except Exception as e:
        envelop.payload["error"] = str(e)
    return envelop
'''


def guess_method(class_name):
    for key, method in METHOD_MAP.items():
        if key in class_name:
            return method
    return "run"


def guess_ctor_args(class_name):
    if any(k in class_name for k in ["Loader", "Reader", "VectorStore", "Retriever"]):
        return "content"
    if any(k in class_name for k in ["Chat", "LLM"]):
        return "**params"
    return ""


def guess_call_args(class_name, method):
    if method in ("load", "load_data"):
        return ""
    return "content"


def should_skip(class_name):
    for pattern in SKIP_PATTERNS:
        if pattern in class_name:
            return True
    return False


def scan_source(source_dir: Path) -> list:
    classes = []
    py_files = list(source_dir.rglob("*.py"))
    total = len(py_files)
    
    for i, f in enumerate(py_files):
        if f.name.startswith("_") or "test" in f.name.lower():
            continue
        try:
            tree = ast.parse(f.read_text(encoding="utf-8"))
            for node in ast.walk(tree):
                if isinstance(node, ast.ClassDef):
                    if node.name.startswith("_"):
                        continue
                    if should_skip(node.name):
                        continue
                    classes.append((node.name, str(f)))
        except:
            pass
        if (i + 1) % 100 == 0:
            print(f"  Scanned {i+1}/{total} files...")
    return classes


def main():
    try:
        import langchain_community
        source = Path(langchain_community.__file__).parent
    except ImportError:
        print("❌ langchain-community not installed!")
        print("   pip install langchain-community")
        return
    
    print("=" * 60)
    print("  LangChain Eater")
    print("=" * 60)
    print(f"📂 Source: {source}")
    
    if OUTPUT_DIR.exists():
        existing = [f for f in OUTPUT_DIR.rglob("*.py") if f.name != "__init__.py"]
        if existing:
            print(f"\n⚠️  {OUTPUT_DIR}/ has {len(existing)} plugins. OVERWRITE!")
    
    print()
    confirm = input("Continue? [y/N] ").strip().lower()
    if confirm not in ("y", "yes"):
        print("Cancelled.")
        return
    
    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)
    for f in OUTPUT_DIR.rglob("*.py"):
        if f.name != "__init__.py":
            f.unlink()
    
    print(f"\n🔍 Scanning...")
    classes = scan_source(source)
    print(f"   {len(classes)} classes found\n")
    
    generated = 0
    written = set()
    
    for class_name, file_path in classes:
        key = class_name.lower()
        if key in written:
            continue
        written.add(key)
        
        try:
            method = guess_method(class_name)
            ctor_args = guess_ctor_args(class_name)
            call_args = guess_call_args(class_name, method)
            
            rel = Path(file_path).relative_to(source)
            module_path = "langchain_community." + str(rel.with_suffix("")).replace("\\", ".").replace("/", ".")
            
            code = PLUGIN_TPL.format(
                description=f"{class_name} from LangChain",
                name=class_name.lower(),
                class_name=class_name,
                module_path=module_path,
                method=method,
                ctor_args=ctor_args,
                call_args=call_args,
            )
            
            (OUTPUT_DIR / f"{class_name.lower()}.py").write_text(code, encoding="utf-8")
            generated += 1
            
            if generated % 100 == 0:
                print(f"  {generated}...")
                
        except Exception as e:
            pass
    
    actual = len([f for f in OUTPUT_DIR.rglob("*.py") if f.name != "__init__.py"])
    print(f"\n✅ {actual} plugins → {OUTPUT_DIR}/")
    print(f"   Restart: python main.py")


if __name__ == "__main__":
    main()