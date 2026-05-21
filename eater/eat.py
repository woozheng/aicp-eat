# eater/eat.py
"""
Universal Eater — Devour ANY Python library into AICP plugins
Usage: python eater/eat.py <library> [source_path]
Examples:
  python eater/eat.py pandas
  python eater/eat.py numpy /path/to/numpy/source
"""

import ast
import sys
from pathlib import Path

# ================================================================
# Config
# ================================================================
METHOD_MAP = {
    "Loader": "load", "Reader": "load_data",
    "Tool": "run", "Wrapper": "run", "QueryRun": "run",
    "Chat": "invoke", "LLM": "invoke", "OpenAI": "invoke",
    "VectorStore": "similarity_search", "Retriever": "get_relevant_documents",
    "Embeddings": "embed_query",
    "DataFrame": "head", "Series": "head", "Index": "tolist",
    "ExcelWriter": "save",
    "Session": "request", "Client": "request",
    "Response": "json", "Request": "prepare",
    "Classifier": "predict", "Regressor": "predict",
    "Forest": "predict", "SVC": "predict",
    "KMeans": "fit_predict", "Scaler": "transform",
}

SKIP_PATTERNS = [
    "Error", "Exception", "Warning",
    "Base", "Meta", "Abstract",
    "Mixin", "Model", "Type", "Ops", "Col",
    "Impl", "Helper", "Util", "Internal",
    "Protocol", "Interface",
    "Dtype", "Grouper", "Agg", "Spec", "Config", "Option", "Param",
    "Descriptor", "Accessor", "Info", "Flags", "Proxy",
    "Factory", "Builder", "Generator",
]

# Only these classes are kept (practical, standalone)
KEEP_CLASSES = [
    "DataFrame", "Series", "Index", "MultiIndex",
    "Categorical", "HDFStore", "ExcelWriter",
]

# ================================================================
# Templates
# ================================================================
CLASS_TPL = '''"""
{description}
Route: POST /api/{bank}/{name}
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
        "route": "/api/{bank}/{name}",
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
        
        try:
            result = instance.{method}({call_args})
        except (AttributeError, TypeError):
            result = instance
        
        if hasattr(result, "to_dict"):
            result = result.to_dict()
        elif hasattr(result, "tolist"):
            result = result.tolist()
        elif hasattr(result, "content"):
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
        envelop.payload["source"] = "{bank}"
    except Exception as e:
        envelop.payload["error"] = str(e)
    return envelop
'''

# FUNC_TPL — 用原始模块路径导入

FUNC_TPL = '''"""
{description}
Route: POST /api/{bank}/{name}
"""

import inspect
from {module_path} import {func_name}

def help():
    try:
        sig = inspect.signature({func_name})
        params = {{}}
        for name, param in sig.parameters.items():
            if param.default is inspect.Parameter.empty:
                params[name] = "<required>"
            else:
                params[name] = str(param.default)
    except:
        params = {{"content": "input"}}
    
    return {{
        "route": "/api/{bank}/{name}",
        "input": {{"params": params}},
        "output": {{"result": "output"}},
        "description": "{description}"
    }}

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        params = payload.get("params", {{}})
        
        if params:
            result = {func_name}(**params)
        else:
            result = {func_name}()
        
        if hasattr(result, "to_dict"):
            result = result.to_dict()
        elif isinstance(result, list):
            result = [str(r) for r in result[:10]]
        else:
            result = str(result)
        
        envelop.payload["result"] = result
        envelop.payload["source"] = "{bank}"
    except Exception as e:
        envelop.payload["error"] = str(e)
    return envelop
'''

# ================================================================
# Helpers
# ================================================================
def guess_method(class_name):
    for key, method in METHOD_MAP.items():
        if key in class_name:
            return method
    return "run"


def guess_ctor_args(class_name):
    if any(k in class_name for k in ["Loader", "Reader", "VectorStore", "Retriever"]):
        return "content"
    if any(k in class_name for k in ["Writer"]):
        return '"output.txt"'
    return ""


def guess_call_args(class_name, method):
    if method in ("load", "load_data", "head", "tail", "describe",
                  "tolist", "to_dict", "save", "render", "prepare", "json"):
        return ""
    return "content"


def should_skip(class_name):
    for pattern in SKIP_PATTERNS:
        if pattern in class_name:
            return True
    return False


def get_exported_names(source_dir: Path) -> set:
    """Get all public names exported in __init__.py."""
    init_file = source_dir / "__init__.py"
    if not init_file.exists():
        init_file = source_dir.parent / "__init__.py"
    if not init_file.exists():
        return set()
    
    try:
        tree = ast.parse(init_file.read_text(encoding="utf-8"))
        exported = set()
        for node in ast.walk(tree):
            if isinstance(node, ast.ImportFrom):
                for alias in node.names:
                    if alias.name != "*":
                        exported.add(alias.name)
            if isinstance(node, ast.Assign):
                for target in node.targets:
                    if isinstance(target, ast.Name) and target.id == "__all__":
                        if isinstance(node.value, (ast.List, ast.Tuple)):
                            for elt in node.value.elts:
                                val = getattr(elt, "value", None) or getattr(elt, "s", None)
                                if val:
                                    exported.add(val)
        return exported
    except:
        return set()


def scan_classes(source_dir: Path) -> list:
    """Extract public standalone classes."""
    classes = []
    py_files = list(source_dir.rglob("*.py"))
    total = len(py_files)
    
    for i, f in enumerate(py_files):
        if f.name.startswith("_") or "test" in f.name.lower():
            continue
        try:
            parts = f.relative_to(source_dir).parts
        except ValueError:
            parts = f.parts
        if any(p.startswith("_") for p in parts):
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


def scan_functions(source_dir: Path) -> list:
    """Extract public standalone functions."""
    funcs = []
    py_files = list(source_dir.rglob("*.py"))
    total = len(py_files)
    
    for i, f in enumerate(py_files):
        if f.name.startswith("_") or "test" in f.name.lower():
            continue
        try:
            parts = f.relative_to(source_dir).parts
        except ValueError:
            parts = f.parts
        if any(p.startswith("_") for p in parts):
            continue
        try:
            tree = ast.parse(f.read_text(encoding="utf-8"))
            for node in ast.walk(tree):
                if isinstance(node, ast.FunctionDef):
                    if node.name.startswith("_"):
                        continue
                    funcs.append((node.name, str(f)))
        except:
            pass
    return funcs


# ================================================================
# Main
# ================================================================
def main():
    if len(sys.argv) < 2:
        print("=" * 60)
        print("  Universal Eater")
        print("=" * 60)
        print()
        print("Usage: python eater/eat.py <library> [source_path]")
        print("Examples:")
        print("  python eater/eat.py pandas")
        print("  python eater/eat.py numpy /path/to/source/")
        print()
        return
    
    LIBRARY = sys.argv[1]
    SOURCE_PATH = sys.argv[2] if len(sys.argv) > 2 else None
    BANK = LIBRARY.lower().replace("-", "_")
    IMPORT_NAME = LIBRARY
    if BANK == "pil":
        IMPORT_NAME = "PIL"
    elif BANK == "sklearn":
        IMPORT_NAME = "sklearn"
    OUTPUT_DIR = Path("plugins") / BANK
    
    print("=" * 60)
    print(f"  Universal Eater — {LIBRARY}")
    print("=" * 60)
    
    if OUTPUT_DIR.exists():
        existing = [f for f in OUTPUT_DIR.rglob("*.py") if f.name != "__init__.py"]
        if existing:
            print(f"\n⚠️  {OUTPUT_DIR}/ has {len(existing)} plugins. OVERWRITE!")
    
    source = None
    if SOURCE_PATH:
        source = Path(SOURCE_PATH)
        if not source.exists():
            print(f"\n❌ Source not found: {source}")
            return
        print(f"📂 Source: {source}")
    else:
        try:
            lib = __import__(BANK)
            source = Path(lib.__file__).parent
            print(f"📂 Installed: {LIBRARY}")
        except ImportError:
            try:
                lib = __import__(LIBRARY)
                source = Path(lib.__file__).parent
                print(f"📂 Installed: {LIBRARY}")
            except ImportError:
                print(f"\n❌ {LIBRARY} not installed!")
                print(f"   pip install {LIBRARY}")
                print(f"   or: python eater/eat.py {LIBRARY} /path/to/source/")
                return
    
    if not source.exists():
        print(f"\n❌ Source not found: {source}")
        return
    
    exported = get_exported_names(source)
    if len(exported) < 10:
        print(f"📋 {len(exported)} names in __all__ (too small, skipping filter)")
        exported = set()
    else:
        print(f"📋 {len(exported)} names in __all__")
    
    print(f"\n📁 Output: {OUTPUT_DIR}/")
    print(f"🔍 Source: {source}")
    print()
    
    confirm = input("Continue? [y/N] ").strip().lower()
    if confirm not in ("y", "yes"):
        print("Cancelled.")
        return
    
    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)
    for f in OUTPUT_DIR.rglob("*.py"):
        if f.name != "__init__.py":
            f.unlink()
    
    generated = 0
    
    # ================================================================
    # Functions first
    # ================================================================
    print(f"\n🔍 Scanning functions...")
    funcs = scan_functions(source)
    print(f"   {len(funcs)} functions found")
    
    if exported:
        funcs = [(n, p) for n, p in funcs if n in exported and not should_skip(n)]
    else:
        funcs = [(n, p) for n, p in funcs if not should_skip(n)]
    print(f"   {len(funcs)} public functions\n")
    
    written = set()
    for func_name, file_path in funcs:
        key = func_name.lower()
        if key in written:
            continue
        written.add(key)
        try:
            # Build module path with original casing
            sub_module = Path(file_path).stem
            if sub_module == "__init__":
                module_path = IMPORT_NAME
            else:
                module_path = f"{IMPORT_NAME}.{sub_module}"
            
            code = FUNC_TPL.format(
                description=f"{func_name} from {LIBRARY}",
                bank=BANK,
                name=func_name.lower(),
                func_name=func_name,
                module_path=module_path,
            )
            (OUTPUT_DIR / f"{func_name.lower()}.py").write_text(code, encoding="utf-8")
            generated += 1
        except Exception as e:
            pass
    # ================================================================
    # Classes second
    # ================================================================
    print(f"🔍 Scanning classes...")
    classes = scan_classes(source)
    print(f"   {len(classes)} classes found")
    
    if KEEP_CLASSES:
        if exported:
            classes = [(n, p) for n, p in classes if n in exported and n in KEEP_CLASSES]
        else:
            classes = [(n, p) for n, p in classes if n in KEEP_CLASSES]
    else:
        if exported:
            classes = [(n, p) for n, p in classes if n in exported]
    print(f"   {len(classes)} practical classes\n")
    
    for class_name, file_path in classes:
        key = class_name.lower()
        if key in written:
            continue
        written.add(key)
        try:
            method = guess_method(class_name)
            ctor_args = guess_ctor_args(class_name)
            call_args = guess_call_args(class_name, method)
            
            code = CLASS_TPL.format(
                description=f"{class_name} from {LIBRARY}",
                bank=BANK,
                name=class_name.lower(),
                class_name=class_name,
                module_path=BANK,
                method=method,
                ctor_args=ctor_args,
                call_args=call_args,
            )
            (OUTPUT_DIR / f"{class_name.lower()}.py").write_text(code, encoding="utf-8")
            generated += 1
        except Exception as e:
            pass
    
    actual = len([f for f in OUTPUT_DIR.rglob("*.py") if f.name != "__init__.py"])
    print(f"\n{'='*60}")
    print(f"✅ {actual} plugins → {OUTPUT_DIR}/")
    print(f"   Restart: python main.py")
    print(f"   Test: curl -X POST http://localhost:9000/api/{BANK}/<name>")


if __name__ == "__main__":
    main()