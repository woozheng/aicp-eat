# gateway/http.py
"""HTTP Gateway — Wraps HTTP requests into Envelops + Static files."""

import asyncio
import json
import uuid
from pathlib import Path
from aiohttp import web

from .. import core                    # 原来是 import core
from ..core import Envelop, engine_route  # 原来是 from core import .


class Gateway:
    def __init__(self, host="0.0.0.0", port=9000, static_dir="www"):
        self.host = host
        self.port = port
        self.static_dir = static_dir
        self.app = web.Application()
        self._request_semaphore = asyncio.Semaphore(100)
        self._runner = None
        self._setup_routes()

    def _setup_routes(self):
        @web.middleware
        async def cors(request, handler):
            if request.method == "OPTIONS":
                return web.Response(status=200, headers={
                    "Access-Control-Allow-Origin": "*",
                    "Access-Control-Allow-Methods": "GET, POST, OPTIONS",
                    "Access-Control-Allow-Headers": "Content-Type, Authorization",
                })
            resp = await handler(request)
            resp.headers["Access-Control-Allow-Origin"] = "*"
            return resp

        self.app.middlewares.append(cors)

        # Health check
        self.app.router.add_get("/health", self._health)

        # API routes
        self.app.router.add_post("/api/{path:.*}", self._handle_api)
        self.app.router.add_get("/api/{path:.*}", self._handle_api_get)

        # Static files — catch all .html files
        self.app.router.add_get("/{filename}.html", self._handle_static)
        self.app.router.add_get("/{filename}.js", self._handle_static)
        self.app.router.add_get("/{filename}.css", self._handle_static)
        self.app.router.add_get("/", self._handle_index)
        self.app.router.add_get("/api/list", self._handle_api_list)
        self.app.router.add_get("/api/pages", self._handle_pages)

    async def _handle_pages(self, request):
        """List all HTML pages in www/ directory."""
        www = Path("www")
        if not www.exists():
            return web.json_response({"pages": []})
        
        pages = []
        for f in sorted(www.rglob("*.html")):
            if f.name == "index.html":  # ← 跳过自己
                continue
            path = str(f.relative_to(www)).replace("\\", "/")
            pages.append({
                "path": path,
                "url": f"/{path}",
                "size": f.stat().st_size,
            })
        
        return web.json_response({"pages": pages})    

    async def _handle_api_list(self, request):
        """List all APIs with help info."""
        apis = []
        for name in core.plugins:
            api_info = {"route": f"/api/{name}"}
            try:
                mod_name = f"plugins.{name.replace('/', '.')}"
                mod = __import__(mod_name, fromlist=["help"])
                if hasattr(mod, "help"):
                    try:
                        help_data = mod.help()
                        # Safe serialize: any non-JSON type → string
                        api_info["help"] = json.loads(
                            json.dumps(help_data, default=str)
                        )
                    except Exception:
                        api_info["help"] = {"error": "help() returned invalid data"}
            except Exception:
                pass
            apis.append(api_info)
        
        return web.json_response({"apis": apis})

        
    async def _handle_static(self, request):
        """Serve static files from www/"""
        filename = request.match_info.get("filename", "index.html")
        # 加上扩展名
        path = request.path.lstrip("/")
        filepath = Path(self.static_dir) / path

        if filepath.exists() and filepath.is_file():
            return web.FileResponse(filepath)

        return web.Response(text="Not found", status=404)

    async def _handle_index(self, request):
        """Serve index.html"""
        filepath = Path(self.static_dir) / "index.html"
        if filepath.exists():
            return web.FileResponse(filepath)
        return web.Response(text="AICP Gateway is running", status=200)

    async def _health(self, request):
        """Health check."""
        try:
            llm_ok = False
            if core.llm:
                try:
                    llm_ok = await core.llm.health_check()
                except Exception:
                    pass

            return web.json_response({
                "status": "ok",
                "agents": len(core.agents),
                "plugins": len(core.plugins),
                "llm": llm_ok,
            })
        except Exception as e:
            return web.json_response({
                "status": "error",
                "error": str(e)[:100]
            })

    async def _handle_api(self, request):
        async with self._request_semaphore:
            path = request.match_info["path"]
            request_id = str(uuid.uuid4())[:8]

            try:
                body = await asyncio.wait_for(request.json(), timeout=10)
            except asyncio.TimeoutError:
                return web.json_response({"error": "Request body read timeout"}, status=408)
            except Exception:
                body = {}

            env = Envelop(
                sender="http",
                receiver=path,
                intent=body.get("intent", "API_CALL") if isinstance(body, dict) else "API_CALL",
                payload=body if isinstance(body, dict) else {"content": str(body)},
                channel_id=f"http:{path}",
            )
            env.message_id = request_id

            print(f"🌐 [{request_id}] POST /api/{path}")
            try:
                result = await asyncio.wait_for(engine_route(env), timeout=120)
                print(f"✅ [{request_id}] Completed")
            except asyncio.TimeoutError:
                print(f"⏱️ [{request_id}] Timeout")
                return web.json_response({"error": "Request timeout"}, status=504)
            except Exception as e:
                print(f"💥 [{request_id}] Error: {e}")
                return web.json_response({"error": str(e)[:200]}, status=500)

            if result:
                return web.json_response(result.payload)
            return web.json_response({"error": "No response"}, status=500)

    async def _handle_api_get(self, request):
        path = request.match_info["path"]
        return web.json_response({
            "route": f"/api/{path}",
            "methods": ["POST"],
            "message": "Use POST to call this API"
        })

    async def start(self):
        try:
            self._runner = web.AppRunner(self.app)
            await self._runner.setup()
            site = web.TCPSite(self._runner, self.host, self.port)
            await site.start()

            print(f"\n{'='*50}")
            print(f"  AICP Gateway")
            print(f"  http://{self.host}:{self.port}")
            print(f"  Agents: {len(core.agents)}")
            print(f"  Plugins: {len(core.plugins)}")
            if self.static_dir and Path(self.static_dir).exists():
                print(f"  Static: {self.static_dir}/")
            print(f"{'='*50}")
            print(f"  Press Ctrl+C to stop\n")

            try:
                while True:
                    await asyncio.sleep(3600)
            except asyncio.CancelledError:
                pass
        except OSError as e:
            if "address already in use" in str(e).lower():
                print(f"\n❌ Port {self.port} is already in use!")
                print(f"   Stop the other process or change port in aicp.yaml")
            else:
                print(f"\n❌ Failed to start: {e}")
        except Exception as e:
            print(f"\n❌ Failed to start: {e}")

    async def stop(self):
        if self._runner:
            await self._runner.cleanup()
            print("   Gateway stopped")