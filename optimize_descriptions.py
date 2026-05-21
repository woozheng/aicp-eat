# optimize_descriptions.py
"""
Local script — Optimize plugin descriptions batch by batch.
Reads plugin code, lets LLM infer function from help() and execute().
Usage: python optimize_descriptions.py <target> [start] [end]
"""

import sys
import re
import asyncio
from pathlib import Path
from openai import AsyncOpenAI

PLUGINS_DIR = Path("plugins")
LLM_CONFIG = {
    "base_url": "https://gpt-agent.cc/v1",
    "api_key": "Your-api_key",
    "model": "kimi-for-coding",
}


async def main():
    if len(sys.argv) < 2:
        print("Usage: python optimize_descriptions.py <target> [start] [end]")
        print("  python optimize_descriptions.py sklearn 0 50")
        sys.exit(1)
    
    target = sys.argv[1]
    start = int(sys.argv[2]) if len(sys.argv) > 2 else 0
    end = int(sys.argv[3]) if len(sys.argv) > 3 else None
    
    scan_dir = PLUGINS_DIR / target
    if not scan_dir.exists():
        print(f"❌ Directory not found: {scan_dir}")
        return
    
    files = [f for f in sorted(scan_dir.rglob("*.py")) 
             if not f.name.startswith("_") and f.name != "__init__.py"]
    
    if end:
        files = files[start:end]
    else:
        files = files[start:]
    
    print(f"🔍 {len(files)} plugins\n")
    
    client = AsyncOpenAI(
        api_key=LLM_CONFIG["api_key"],
        base_url=LLM_CONFIG["base_url"],
        timeout=30,
    )
    
    optimized = 0
    for i, f in enumerate(files):
        code = f.read_text(encoding="utf-8")
        
        if "def help():" not in code:
            continue
        
        match = re.search(r'"description":\s*"([^"]*)"', code)
        if not match:
            continue
        
        old_desc = match.group(1)
        route_path = str(f.relative_to(PLUGINS_DIR)).replace("\\", "/").replace(".py", "")
        
        # Extract help() for context
        help_match = re.search(r'def help\(\):.*?return\s*(\{.*?\})\s*$', code, re.DOTALL | re.MULTILINE)
        help_info = help_match.group(1)[:300] if help_match else "{}"
        print(f"Route: {route_path}")
        print(f"Old: {old_desc}")
        try:
           
            resp = await client.chat.completions.create(
    model=LLM_CONFIG["model"],
    messages=[
        {"role": "system", "content": """You are optimizing AICP plugin descriptions. 
Look at the route name and infer what this Python function/class does.
Examples:
- requests/get → "Make an HTTP GET request to a URL"
- requests/session → "Maintains persistent cookies and headers across multiple HTTP requests"  
- pandas/read_csv → "Read a CSV file into a DataFrame"
- numpy/mean → "Compute the arithmetic mean of array elements"
- sklearn/randomforestclassifier → "Random forest classifier for classification tasks"

Now write a 5-10 word description for this plugin:"""},
    {"role": "user", "content": f"Route: {route_path}\nOld: {old_desc}\n\nBetter description:"}
    ],
    max_tokens=30,
    temperature=0.3,
)
             
            print(f"RAW: '{resp.choices[0].message.content}'")
            new_desc = resp.choices[0].message.content.strip().strip('"').strip("'")
            print(f"-----{new_desc}")
            if new_desc and new_desc != old_desc and len(new_desc) > 5:
                new_code = code.replace(
                    f'"description": "{old_desc}"',
                    f'"description": "{new_desc}"'
                )
                f.write_text(new_code, encoding="utf-8")
                print(f"  ✅ {route_path}")
                print(f"     {old_desc}")
                print(f"     → {new_desc}\n")
                optimized += 1
            else:
                print(f"  ⏭️  {new_desc} (unchanged)\n")
        
        except Exception as e:
            print(f"  ❌ {route_path}: {e}\n")
        
        await asyncio.sleep(0.05)
    
    print(f"✅ {optimized} optimized, {len(files) - optimized} unchanged")


if __name__ == "__main__":
    asyncio.run(main())