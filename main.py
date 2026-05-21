# main.py
"""AICP Engine — Startup with graceful shutdown."""

import asyncio
import shutil
import importlib.util
from pathlib import Path

from . import core 
from .core import Bus
from .llm import LLM
from .config import load_config
from .gateway.http import Gateway

PACKAGE_DIR = Path(__file__).parent


async def run(config_path: str = None):
    print("=" * 50)
    print("  AICP — Information Field Protocol Engine")
    print("=" * 50)

    # Auto-init directories and default files
    Path("data").mkdir(exist_ok=True)
    
    # 拷贝 aicp.yaml（这个需要，因为用户要改配置）
    config_file = Path("aicp.yaml")
    if not config_file.exists():
        src = Path(__file__).parent / "aicp.yaml"
        if src.exists():
            shutil.copy(src, config_file)
    
    # www/ 和 plugins/ 已经在包安装目录里了，Gateway 直接从当前目录读
    # 如果用户目录没有，从包目录拷贝
    for d in ["www", "plugins"]:
        if not Path(d).exists():
            src = Path(__file__).parent / d
            if src.exists():
                shutil.copytree(src, d)

    # 1. Load config
    print("\n📋 Loading config...")
    config = load_config(config_path)
    print(f"   Host: {config['host']}:{config['port']}")
    print(f"   Plugins dir: {config['plugins_dir']}")

    # 2. Init LLM
    print("\n🔧 Initializing LLM...")
    llm = LLM(config)
    core.llm = llm

    # 3. Init Bus
    print("\n📡 Initializing Bus...")
    core.bus = Bus()

    # 4. Load plugins
    print("\n📦 Loading plugins...")
    load_plugins(config["plugins_dir"])

    # 5. Load agents
    print("\n👤 Loading agents...")
    load_agents(config)

    # 6. Start Gateway
    print()
    gateway = Gateway(
        host=config["host"],
        port=config["port"],
        static_dir=config.get("static_dir", "www"),
    )
    await gateway.start()


def load_plugins(plugins_dir: str = "plugins"):
    """Auto-scan and register all plugins."""
    plugin_path = Path(plugins_dir)
    if not plugin_path.exists():
        print("   (no plugins directory)")
        return

    loaded = 0
    errors = 0
    for f in sorted(plugin_path.rglob("*.py")):
        if f.name.startswith("_"):
            continue
        if f.parent.name == "__pycache__":
            continue

        name = str(f.relative_to(plugin_path)).replace("\\", "/").replace(".py", "")

        spec = importlib.util.spec_from_file_location(
            f"plugin_{name.replace('/', '_')}", f
        )
        if not spec:
            continue

        mod = importlib.util.module_from_spec(spec)
        try:
            spec.loader.exec_module(mod)
        except Exception:
            errors += 1
            continue

        if hasattr(mod, "execute"):
            core.plugins[name] = mod.execute
            loaded += 1
            print(f"  ✅ /api/{name}")

    print(f"   {loaded} loaded, {errors} skipped")


def load_agents(config: dict):
    """Register agents from config."""
    robots = config.get("robots", [])
    for bot in robots:
        agent_id = bot.get("id", bot.get("user_id", ""))
        if not agent_id:
            continue
        core.agents[agent_id] = {
            "id": agent_id,
            "display_name": bot.get("display_name", agent_id),
            "workflow": bot.get("workflow", bot.get("action", [])),
            "prompts": bot.get("prompts", ""),
            "brain": bot.get("brain", ""),
        }
        print(f"  🤖 {agent_id}")

    if not robots:
        print("   (no agents configured)")

    groups = config.get("groups", [])
    for group in groups:
        group_id = group.get("id", "")
        members = group.get("members", [])
        if group_id and members:
            channel = f"grp.{group_id}"
            for member in members:
                core.bus.subscribe(member, channel)
            print(f"  👥 {group_id}: {len(members)} members")


if __name__ == "__main__":
    asyncio.run(run())