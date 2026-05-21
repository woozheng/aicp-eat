"""AICP 命令行入口"""

import sys
import asyncio


def main():
    config_path = None

    # 简单解析命令行参数
    args = sys.argv[1:]
    for i, arg in enumerate(args):
        if arg == "--config" and i + 1 < len(args):
            config_path = args[i + 1]
        elif arg in ("--help", "-h"):
            print("AICP — Eat Every Python Library → HTTP API")
            print("Usage: aicp [--config PATH]")
            print("Open http://localhost:9000 after startup")
            return

    # 直接调你 main.py 里的 run()
    from .main import run
    asyncio.run(run(config_path=config_path))