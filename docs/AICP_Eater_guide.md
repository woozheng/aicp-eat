# AICP Eater Guide
## Rule
Eater script scans a Python library and generates one AICP plugin per class.

```python
# eater/eat_xxx.py — run once, generate all plugins
```
### Eater Template
```python
"""
{Target} Eater — Devour {library} into AICP plugins
Output: plugins/{library}/
"""

import ast
from pathlib import Path

OUTPUT_DIR = Path("plugins") / "{library}"

METHOD_MAP = {
    "Loader": "load", "Tool": "run", "Chat": "invoke",
    "VectorStore": "similarity_search", "Classifier": "predict"
}

PLUGIN_TPL = '''"""
{description}
Route: POST /api/{bank}/{name}
"""

from {module_path} import {class_name}

def help():
    return {{
        "route": "/api/{bank}/{name}",
        "input": {{"content": "input"}},
        "output": {{"result": "output"}},
        "description": "{description}"
    }}

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        content = payload.get("content", "")
        params = payload.get("params", {{}})
        
        try:
            instance = {class_name}({ctor_args})
        except:
            instance = {class_name}()
        
        result = instance.{method}({call_args})
        
        if hasattr(result, "content"):
            result = result.content
        elif isinstance(result, list):
            result = [str(r) for r in result[:10]]
        
        envelop.payload["result"] = str(result)
        envelop.payload["source"] = "{bank}"
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
    if "Loader" in class_name or "VectorStore" in class_name:
        return "content"
    return ""


def guess_call_args(class_name, method):
    if method in ("load", "load_data"):
        return ""
    return "content"


def scan_source(source_dir: Path):
    """Scan all .py files and extract public classes."""
    classes = []
    for f in source_dir.rglob("*.py"):
        if f.name.startswith("_") or "test" in f.name.lower():
            continue
        try:
            tree = ast.parse(f.read_text(encoding="utf-8"))
            for node in ast.walk(tree):
                if isinstance(node, ast.ClassDef) and not node.name.startswith("_"):
                    classes.append((node.name, str(f)))
        except:
            pass
    return classes


def generate_plugins():
    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)
    
    # Clean old
    for f in OUTPUT_DIR.rglob("*.py"):
        if f.name != "__init__.py":
            f.unlink()
    
    # Scan
    source = Path("temp/{library}")  # or pip show {library} location
    if not source.exists():
        import subprocess
        subprocess.run(["pip", "install", "{library}"], capture_output=True)
        import {library}
        source = Path({library}.__file__).parent
    
    classes = scan_source(source)
    print(f"Found {{len(classes)}} classes")
    
    # Generate
    for class_name, file_path in classes:
        try:
            method = guess_method(class_name)
            ctor_args = guess_ctor_args(class_name)
            call_args = guess_call_args(class_name, method)
            
            # Build module path from file path
            rel = Path(file_path).relative_to(source.parent.parent)
            module_path = str(rel.with_suffix("")).replace("/", ".").replace("\\", ".")
            
            code = PLUGIN_TPL.format(
                description=f"{{class_name}} from {{library}}",
                bank="{library}",
                name=class_name.lower(),
                class_name=class_name,
                module_path=module_path,
                method=method,
                ctor_args=ctor_args,
                call_args=call_args,
            )
            
            out_file = OUTPUT_DIR / f"{{class_name.lower()}}.py"
            out_file.write_text(code, encoding="utf-8")
            print(f"  ✅ {{class_name}}")
            
        except Exception as e:
            print(f"  ❌ {{class_name}}: {{e}}")
    
    print(f"\\n✅ Generated plugins in {{OUTPUT_DIR}}")


if __name__ == "__main__":
    generate_plugins()
```    
## How AI Generates Eater

Copy the template above

Replace {library} with the target library name

Replace {Target} with the library display name

Replace {bank} with library lower case

Fill METHOD_MAP with library-specific method patterns

### Output
```text
plugins/{library}/
├── class1.py    → POST /api/{library}/class1
├── class2.py    → POST /api/{library}/class2
└── ...
```

Each plugin has help() and execute(envelop, agent).

Copy template. Replace placeholders. Run. Hundreds of APIs generated.