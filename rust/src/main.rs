use axum::{
    extract::Path,
    routing::{get, post},
    Json, Router,
};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::fs;
use std::path::Path as StdPath;
use std::process::Command;
use syn::visit::Visit;

// ============================================================
// AICP Protocol v3.0 — Envelop
// ============================================================

#[derive(Debug, Serialize, Deserialize, Clone)]
struct Envelop {
    sender: String,
    receiver: String,
    intent: String,
    payload: HashMap<String, serde_json::Value>,
    trace_id: String,
    message_id: String,
    channel_id: String,
    ttl: i32,
    meta: HashMap<String, serde_json::Value>,
}

// ============================================================
// Plugin trait
// ============================================================

#[async_trait::async_trait]
trait Plugin {
    fn name(&self) -> &str;
    async fn execute(&self, envelop: Envelop) -> Option<Envelop>;
}

// ============================================================
// Engine — Blind Router
// ============================================================

struct Engine {
    routes: HashMap<String, Vec<Box<dyn Plugin + Send + Sync>>>,
}

impl Engine {
    fn new() -> Self {
        Engine {
            routes: HashMap::new(),
        }
    }

    fn register(&mut self, id: String, plugins: Vec<Box<dyn Plugin + Send + Sync>>) {
        self.routes.insert(id, plugins);
    }

    async fn route(&self, mut envelop: Envelop) -> Option<Envelop> {
        if envelop.ttl <= 0 {
            return None;
        }
        envelop.ttl -= 1;

        let original_receiver = envelop.receiver.clone();
        let plugins = self.routes.get(&original_receiver)?;

        if envelop.sender == original_receiver {
            return None;
        }

        envelop.receiver = String::new();
        let mut current = envelop;

        for plugin in plugins {
            match plugin.execute(current).await {
                Some(env) => current = env,
                None => return None,
            }
        }

        current.sender = original_receiver;
        Some(current)
    }
}

// ============================================================
// Function Scanner
// ============================================================

#[derive(Debug, Serialize, Deserialize, Clone)]
struct FuncInfo {
    name: String,
    params: Vec<ParamInfo>,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
struct ParamInfo {
    name: String,
    #[serde(rename = "type")]
    ty: String,
}

struct FuncVisitor {
    functions: Vec<FuncInfo>,
}

impl<'ast> Visit<'ast> for FuncVisitor {
    fn visit_item_fn(&mut self, f: &'ast syn::ItemFn) {
        let name = f.sig.ident.to_string();
        let params: Vec<ParamInfo> = f
            .sig
            .inputs
            .iter()
            .filter_map(|arg| {
                if let syn::FnArg::Typed(pat) = arg {
                    Some(ParamInfo {
                        name: quote::quote!(#pat.pat).to_string(),
                        ty: quote::quote!(#pat.ty).to_string(),
                    })
                } else {
                    None
                }
            })
            .collect();
        self.functions.push(FuncInfo { name, params });
        syn::visit::visit_item_fn(self, f);
    }
}

fn eat_crate(crate_name: &str) -> Result<Vec<FuncInfo>, Box<dyn std::error::Error>> {
    println!("Scanning {} from local cache...", crate_name);

    // 确保 crate 已下载
    let _ = Command::new("cargo").args(["add", crate_name]).output();

    // 从 cargo 缓存找
    let home = std::env::var("USERPROFILE").unwrap_or_else(|_| ".".to_string());
    let cache_dir = format!("{}\\.cargo\\registry\\src", home);

    if let Some(dir) = find_crate_dir(StdPath::new(&cache_dir), crate_name) {
        println!("Found at: {}", dir);
        return scan_directory(&dir);
    }

    println!("Not found in cache, scanning current project...");
    scan_directory("./src")
}

fn find_crate_dir(base: &StdPath, name: &str) -> Option<String> {
    if let Ok(entries) = fs::read_dir(base) {
        for e in entries.flatten() {
            let p = e.path();
            if p.is_dir() {
                if let Ok(sub) = fs::read_dir(&p) {
                    for se in sub.flatten() {
                        let sp = se.path();
                        let dir_name = sp.file_name().unwrap().to_string_lossy();
                        if dir_name.contains(name) && sp.join("src").exists() {
                            return Some(sp.to_string_lossy().to_string());
                        }
                    }
                }
            }
        }
    }
    None
}

fn scan_directory(dir: &str) -> Result<Vec<FuncInfo>, Box<dyn std::error::Error>> {
    let mut all = Vec::new();
    walk_dir(StdPath::new(dir), &mut all);
    Ok(all)
}

fn walk_dir(dir: &StdPath, functions: &mut Vec<FuncInfo>) {
    if let Ok(entries) = fs::read_dir(dir) {
        for e in entries.flatten() {
            let p = e.path();
            if p.is_dir() {
                walk_dir(&p, functions);
            } else if p.extension().map_or(false, |x| x == "rs") {
                if let Ok(content) = fs::read_to_string(&p) {
                    if let Ok(file) = syn::parse_file(&content) {
                        let mut v = FuncVisitor {
                            functions: Vec::new(),
                        };
                        v.visit_file(&file);
                        functions.extend(v.functions);
                    }
                }
            }
        }
    }
}

// ============================================================
// HTTP Server
// ============================================================

#[tokio::main]
async fn main() {
    let args: Vec<String> = std::env::args().collect();

    // ----- Eat mode -----
    if args.len() >= 3 && args[1] == "eat" {
        let name = &args[2];
        println!("Eating {}...", name);

        match eat_crate(name) {
            Ok(fns) => {
                println!("Found {} functions", fns.len());
                let json = serde_json::to_string_pretty(&fns).unwrap();
                let filename = format!("plugins/{}.json", name.replace('/', "_"));
                fs::create_dir_all("plugins").unwrap();
                fs::write(&filename, json).unwrap();
                println!("Saved to {}", filename);
            }
            Err(e) => eprintln!("Error: {}", e),
        }
        return;
    }

    // ----- Server mode -----
    println!("AICP Rust Engine: http://localhost:9000");

    let app = Router::new()
        .route("/", get(home))
         .route("/test", get(|| async { "test ok" }))
        .route("/plugin/{name}", get(plugin_detail))
        .route("/api/{name}", post(api_handler));
println!("Routes registered: /api/{{name}}");
    let listener = tokio::net::TcpListener::bind("0.0.0.0:9000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

// ----- Home page -----
async fn home() -> axum::response::Html<String> {
    let mut plugins = Vec::new();
    if let Ok(entries) = fs::read_dir("plugins") {
        for e in entries.flatten() {
            let n = e.file_name().to_string_lossy().to_string();
            if n.ends_with(".json") {
                plugins.push(n.trim_end_matches(".json").to_string());
            }
        }
    }

    let mut html = String::from(
        r#"<!DOCTYPE html>
<html>
<head>
    <title>AICP Rust Engine</title>
    <style>
        body { font-family: sans-serif; background: #0d1117; color: #c9d1d9; padding: 40px; }
        h1 { color: #58a6ff; }
        .plugin { background: #161b22; border: 1px solid #30363d; border-radius: 6px; padding: 16px; margin: 10px 0; }
        .plugin a { color: #7ee787; font-size: 18px; text-decoration: none; }
    </style>
</head>
<body>
    <h1>🦀 AICP Rust Engine</h1>
    <p>Eat Rust crates. Serve via HTTP. Just curl.</p>
"#,
    );

    for p in &plugins {
        html.push_str(&format!(
            r#"<div class="plugin"><a href="/plugin/{}">{}</a></div>"#,
            p, p
        ));
    }

    html.push_str("</body></html>");
    axum::response::Html(html)
}

// ----- Plugin detail page -----
async fn plugin_detail(Path(name): Path<String>) -> axum::response::Html<String> {
    let filepath = format!("plugins/{}.json", name);
    let content = fs::read_to_string(&filepath).unwrap_or_default();
    let fns: Vec<FuncInfo> = serde_json::from_str(&content).unwrap_or_default();

    let mut html = format!(
        r#"<!DOCTYPE html>
<html>
<head>
    <title>{0}</title>
    <style>
        body {{ font-family: sans-serif; background: #0d1117; color: #c9d1d9; padding: 20px; }}
        h1 {{ color: #58a6ff; }}
        .back {{ margin-bottom: 20px; }} .back a {{ color: #58a6ff; }}
        .fn {{ background: #161b22; border: 1px solid #30363d; padding: 12px; margin: 8px 0; border-radius: 6px; cursor: pointer; }}
        .fn:hover {{ border-color: #58a6ff; }}
        .name {{ color: #d2a8ff; font-weight: bold; }}
        .params {{ color: #8b949e; font-size: 13px; margin-top: 4px; }}
        .search {{ margin-bottom: 20px; }}
        .search input {{ background: #161b22; border: 1px solid #30363d; color: #c9d1d9; padding: 10px; width: 300px; border-radius: 6px; }}
        #result {{ margin-top: 20px; padding: 15px; background: #161b22; border-radius: 6px; display: none; }}
        #result pre {{ color: #7ee787; margin: 0; }}
    </style>
</head>
<body>
    <p class="back"><a href="/">← Back</a></p>
    <h1>{0}</h1>
    <p>{1} functions</p>
    <div class="search"><input type="text" placeholder="Filter functions..." oninput="filter(this.value)"></div>
    <div id="funcs">
"#,
        name,
        fns.len()
    );

    for f in &fns {
        let ps: Vec<String> = f
            .params
            .iter()
            .map(|p| format!("{}: {}", p.name, p.ty))
            .collect();
        html.push_str(&format!(
            r#"<div class="fn" onclick="call('{}')">
                <div class="name">{}</div>
                <div class="params">{}</div>
            </div>"#,
            f.name, f.name, ps.join(", ")
        ));
    }

    html.push_str(&format!(
        r#"</div>
<div id="result"></div>
<script>
function filter(q) {{
    document.querySelectorAll('.fn').forEach(f => {{
        f.style.display = f.querySelector('.name').textContent.toLowerCase().includes(q.toLowerCase()) ? '' : 'none';
    }});
}}
async function call(func) {{
    const r = document.getElementById('result');
    r.style.display = 'block';
    r.innerHTML = 'Calling <b>{}</b>...';
    try {{
        const resp = await fetch('/api/{}', {{
            method: 'POST',
            headers: {{'Content-Type': 'application/json'}},
            body: JSON.stringify({{meta:{{function: func}}, payload:{{args:{{}}}}}})
        }});
        const data = await resp.json();
        r.innerHTML = '<pre>' + JSON.stringify(data, null, 2) + '</pre>';
    }} catch(e) {{
        r.innerHTML = '<pre>Error: ' + e.message + '</pre>';
    }}
}}
</script>
</body></html>"#,
        name, name
    ));

    axum::response::Html(html)
}

// ----- API handler -----
async fn api_handler(
    Path(name): Path<String>,
    Json(body): Json<HashMap<String, serde_json::Value>>,
) -> Json<HashMap<String, serde_json::Value>> {
    let filepath = format!("plugins/{}.json", name);
    let content = fs::read_to_string(&filepath).unwrap_or_default();
    let fns: Vec<FuncInfo> = serde_json::from_str(&content).unwrap_or_default();

    let func_name = body
        .get("meta")
        .and_then(|m| m.get("function"))
        .and_then(|f| f.as_str())
        .unwrap_or("");

    let found = fns.iter().any(|f| f.name == func_name);

    let mut result = HashMap::new();
    if found {
        result.insert("result".to_string(), serde_json::Value::String(format!("{}() called successfully", func_name)));
        result.insert("source".to_string(), serde_json::Value::String(name));
    } else {
        result.insert("error".to_string(), serde_json::Value::String(format!("function {} not found in {}", func_name, name)));
    }

    Json(result)
}