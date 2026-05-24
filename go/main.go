package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"net/http"
	"os"
	"os/exec"
	
	"regexp"
	"strings"
	"sync"

	"golang.org/x/tools/go/packages"
)

// AICP Envelop
type Envelop struct {
	Sender    string                 `json:"sender"`
	Receiver  string                 `json:"receiver"`
	Intent    string                 `json:"intent"`
	Payload   map[string]interface{} `json:"payload"`
	TraceID   string                 `json:"trace_id"`
	MessageID string                 `json:"message_id"`
	TTL       int                    `json:"ttl"`
	Meta      map[string]interface{} `json:"meta"`
}

// 动态包注册表 - 从 plugins/ 目录自动加载
var (
	pkgRegistry = make(map[string]interface{})
	registryMu  sync.RWMutex
)

// registerPlugin 从插件源码提取包路径并尝试加载
func registerPlugin(pluginName string) {
	filePath := fmt.Sprintf("plugins/%s.go", pluginName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	// 提取 import 路径
	reImport := regexp.MustCompile(`"([^"]+)"`)
	lines := strings.Split(string(content), "\n")
	var importPath string
	for _, line := range lines {
		if match := reImport.FindStringSubmatch(line); len(match) > 1 {
			candidate := match[1]
			if strings.Contains(candidate, ".") &&
				!strings.Contains(candidate, "encoding/") &&
				!strings.Contains(candidate, "net/") &&
				!strings.Contains(candidate, "fmt") &&
				!strings.Contains(candidate, "reflect") &&
				!strings.Contains(candidate, "sync") {
				importPath = candidate
				break
			}
		}
	}

	if importPath == "" {
		return
	}

	// 从插件源码提取 shortName（类型名）
	// 找到 var Funcs 后面的 import 块里的包路径
	shortName := extractShortName(string(content), importPath)
	if shortName == "" {
		return
	}

	// 尝试动态加载包
	pkg, err := loadPackage(importPath, shortName)
	if err != nil {
		fmt.Printf("  [warn] cannot load %s: %v\n", importPath, err)
		return
	}

	registryMu.Lock()
	pkgRegistry[importPath] = pkg
	registryMu.Unlock()
	fmt.Printf("  [registered] %s -> %s\n", pluginName, importPath)
}

// extractShortName 从插件源码提取包别名
func extractShortName(content, importPath string) string {
	// 从 Funcs 的 JSON 里找第一个函数，然后从 Execute 里找类型引用
	// 更简单的方法：直接用正则找 import 行的别名
	re := regexp.MustCompile(`"` + regexp.QuoteMeta(importPath) + `"`)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if re.MatchString(line) {
			// 检查是否有别名
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "\"") || strings.HasPrefix(trimmed, "_") {
				// 没有别名，用包路径最后一段
				parts := strings.Split(importPath, "/")
				return parts[len(parts)-1]
			}
			// 有别名，取别名
			parts := strings.Fields(trimmed)
			if len(parts) > 0 && parts[0] != "\"" {
				return strings.TrimSpace(parts[0])
			}
		}
	}
	parts := strings.Split(importPath, "/")
	return parts[len(parts)-1]
}

// loadPackage 尝试加载包并返回其 reflect.Value
func loadPackage(importPath, shortName string) (interface{}, error) {
	// 用 go/packages 加载包信息
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, importPath)
	if err != nil || len(pkgs) == 0 {
		return nil, fmt.Errorf("packages.Load failed: %v", err)
	}

	pkg := pkgs[0]
	if pkg.Types == nil {
		return nil, fmt.Errorf("no types info")
	}

	// 返回包的类型信息，后续用 types 反射
	return pkg.Types, nil
}

// callFunction 反射调用函数
func callFunction(pkgTypes interface{}, funcName string) (string, error) {
	pkg := pkgTypes.(*types.Package)
	scope := pkg.Scope()

	obj := scope.Lookup(funcName)
	if obj == nil {
		return "", fmt.Errorf("function %s not found", funcName)
	}

	// 对于包级函数，直接返回函数名作为结果
	// 真正执行需要编译时链接，这里做 demo 展示
	return fmt.Sprintf("%s() called from %s", funcName, pkg.Path()), nil
}

// EatGo 吃掉 Go 包
func EatGo(pkgPath string) error {
	fmt.Printf("Downloading %s...\n", pkgPath)
	exec.Command("go", "get", pkgPath).Run()

	fmt.Printf("Scanning %s...\n", pkgPath)

	cfg := &packages.Config{
		Mode: packages.NeedFiles | packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, pkgPath)
	if err != nil {
		return fmt.Errorf("failed to load package: %v", err)
	}

	functions := []map[string]interface{}{}
	fset := token.NewFileSet()

	for _, pkg := range pkgs {
		for _, file := range pkg.GoFiles {
			f, err := parser.ParseFile(fset, file, nil, 0)
			if err != nil {
				continue
			}
			for _, decl := range f.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok {
					if fn.Name.IsExported() {
						params := []map[string]string{}
						if fn.Type.Params != nil {
							for _, p := range fn.Type.Params.List {
								paramType := types.ExprString(p.Type)
								if len(p.Names) > 0 {
									for _, name := range p.Names {
										params = append(params, map[string]string{
											"name": name.Name,
											"type": paramType,
										})
									}
								} else {
									params = append(params, map[string]string{
										"name": "",
										"type": paramType,
									})
								}
							}
						}
						functions = append(functions, map[string]interface{}{
							"name":   fn.Name.Name,
							"params": params,
						})
					}
				}
			}
		}
	}

	if len(functions) == 0 {
		return fmt.Errorf("no exported functions found in %s", pkgPath)
	}

	pkgName := strings.ReplaceAll(pkgPath, "/", "_")
	pkgName = strings.ReplaceAll(pkgName, ".", "_")

	code := generatePlugin(pkgPath, pkgName, functions)

	os.MkdirAll("plugins", 0755)
	filename := fmt.Sprintf("plugins/%s.go", pkgName)
	os.WriteFile(filename, []byte(code), 0644)

	fmt.Printf("Eaten: %s -> %d functions -> %s\n", pkgPath, len(functions), filename)
	return nil
}

// 生成 Plugin 代码
func generatePlugin(pkgPath, pkgName string, functions []map[string]interface{}) string {
	funcListJSON, _ := json.MarshalIndent(functions, "", "  ")

	return fmt.Sprintf(`package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"%s"
)

var Funcs = %s

func GetHelp() map[string]interface{} {
	var parsed []map[string]interface{}
	json.Unmarshal([]byte(Funcs), &parsed)
	return map[string]interface{}{
		"route":     "/api/%s",
		"package":   "%s",
		"total":     len(parsed),
		"functions": parsed,
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var env map[string]interface{}
	json.NewDecoder(r.Body).Decode(&env)
	meta := env["meta"].(map[string]interface{})
	funcName := meta["function"].(string)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, ` + "`" + `{"result":"%%s() called","source":"%s"}` + "`" + `, funcName)
}
`, pkgPath, string(funcListJSON), pkgName, pkgPath, pkgPath)
}

// 启动服务
func startServer() {
	// 自动注册所有插件
	fmt.Println("Loading plugins...")
	entries, _ := os.ReadDir("plugins")
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".go") {
			pluginName := strings.TrimSuffix(e.Name(), ".go")
			registerPlugin(pluginName)
		}
	}
	fmt.Printf("Registered %d packages\n\n", len(pkgRegistry))

	// 首页
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		entries, _ := os.ReadDir("plugins")
		var plugins []string
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".go") {
				name := strings.TrimSuffix(e.Name(), ".go")
				plugins = append(plugins, name)
			}
		}
		pluginsJSON, _ := json.Marshal(plugins)
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>AICP Go Engine</title>
    <style>
        body { font-family: 'Segoe UI', sans-serif; background: #0d1117; color: #c9d1d9; padding: 40px; }
        h1 { color: #58a6ff; }
        .sub { color: #8b949e; margin-bottom: 30px; }
        .plugin { background: #161b22; border: 1px solid #30363d; border-radius: 6px; padding: 16px; margin: 10px 0; }
        .plugin a { color: #7ee787; font-size: 18px; text-decoration: none; }
    </style>
</head>
<body>
    <h1>AICP Go Engine</h1>
    <p class="sub">Eat Go packages. Serve via HTTP. Just curl.</p>
    <div id="list">Loading...</div>
    <script>
        const plugins = %s;
        let html = '';
        plugins.forEach(p => {
            html += '<div class="plugin"><a href="/plugin/' + p + '">' + p + '</a></div>';
        });
        document.getElementById('list').innerHTML = html || '<p>No packages eaten yet.</p>';
    </script>
</body>
</html>`, string(pluginsJSON))
	})

	// 插件详情页
	http.HandleFunc("/plugin/", func(w http.ResponseWriter, r *http.Request) {
		pluginName := strings.TrimPrefix(r.URL.Path, "/plugin/")
		filePath := fmt.Sprintf("plugins/%s.go", pluginName)

		content, err := os.ReadFile(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		contentStr := string(content)
		re := regexp.MustCompile(`var\s+(?:Funcs|funcs)\s*=\s*(\[[\s\S]*?\n\])`)
		matches := re.FindStringSubmatch(contentStr)
		if len(matches) < 2 {
			http.Error(w, "funcs JSON not found", 500)
			return
		}

		funcsJSON := matches[1]
		var functions []map[string]interface{}
		if err := json.Unmarshal([]byte(funcsJSON), &functions); err != nil {
			http.Error(w, fmt.Sprintf("JSON parse error: %v", err), 500)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<html><head><title>%s</title>
<style>
body { font-family: 'Segoe UI', sans-serif; background: #0d1117; color: #c9d1d9; padding: 20px; }
h1 { color: #58a6ff; }
.back { color: #8b949e; margin-bottom: 20px; }
.back a { color: #58a6ff; }
.fn { background: #161b22; border: 1px solid #30363d; padding: 12px; margin: 8px 0; border-radius: 6px; cursor: pointer; }
.fn:hover { border-color: #58a6ff; }
.name { color: #d2a8ff; font-weight: bold; }
.params { color: #8b949e; font-size: 13px; margin-top: 4px; }
.search { margin-bottom: 20px; }
.search input { background: #161b22; border: 1px solid #30363d; color: #c9d1d9; padding: 10px; width: 300px; border-radius: 6px; font-size: 14px; }
</style></head><body>
<p class="back"><a href="/">← Back</a></p>
<h1>%s</h1>
<p>%d functions</p>
<div class="search"><input type="text" placeholder="Filter functions..." oninput="filter(this.value)"></div>
<div id="funcs">`, pluginName, pluginName, len(functions))

		for _, fn := range functions {
			name := fn["name"].(string)
			params := fn["params"].([]interface{})
			paramsStr := ""
			for i, p := range params {
				if i > 0 {
					paramsStr += ", "
				}
				pm := p.(map[string]interface{})
				paramsStr += fmt.Sprintf("%s: %s", pm["name"], pm["type"])
			}
			fmt.Fprintf(w, `<div class="fn" data-name="%s" onclick="callFunction('%s','%s')">
<div class="name">%s</div>
<div class="params">%s</div>
</div>`, name, pluginName, name, name, paramsStr)
		}

		fmt.Fprint(w, `</div>
<div id="result" style="margin-top:20px;padding:15px;background:#161b22;border-radius:6px;display:none;"></div>
<script>
function filter(query) {
    document.querySelectorAll('.fn').forEach(fn => {
        fn.style.display = fn.dataset.name.toLowerCase().includes(query.toLowerCase()) ? '' : 'none';
    });
}
function callFunction(plugin, func) {
    fetch('/api/' + plugin, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({meta:{function:func},payload:{args:{}}})
    })
    .then(r => r.json())
    .then(data => {
        const r = document.getElementById('result');
        r.style.display = 'block';
        r.innerHTML = '<strong>Result:</strong> <pre>' + JSON.stringify(data, null, 2) + '</pre>';
    });
}
</script>
</body></html>`)
	})

	// API 路由
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		pluginName := strings.TrimPrefix(r.URL.Path, "/api/")
		filePath := fmt.Sprintf("plugins/%s.go", pluginName)

		content, err := os.ReadFile(filePath)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"error":"plugin not found"}`)
			return
		}

		reImport := regexp.MustCompile(`"([^"]+)"`)
		lines := strings.Split(string(content), "\n")
		var importPath string
		for _, line := range lines {
			if match := reImport.FindStringSubmatch(line); len(match) > 1 {
				candidate := match[1]
				if strings.Contains(candidate, ".") &&
					!strings.Contains(candidate, "encoding/") &&
					!strings.Contains(candidate, "net/") &&
					!strings.Contains(candidate, "fmt") &&
					!strings.Contains(candidate, "reflect") &&
					!strings.Contains(candidate, "sync") {
					importPath = candidate
					break
				}
			}
		}

		if importPath == "" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"error":"import not found"}`)
			return
		}

		registryMu.RLock()
		pkg, ok := pkgRegistry[importPath]
		registryMu.RUnlock()

		if !ok {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"error":"package %s not registered. Did you restart after eating?"}`, importPath)
			return
		}

		var env map[string]interface{}
		json.NewDecoder(r.Body).Decode(&env)
		meta := env["meta"].(map[string]interface{})
		funcName := meta["function"].(string)

		result, err := callFunction(pkg, funcName)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			fmt.Fprintf(w, `{"error":"%v"}`, err)
		} else {
			fmt.Fprintf(w, `{"result":"%s","source":"%s"}`, result, importPath)
		}
	})

	fmt.Println("\nAICP Go Engine: http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func main() {
	if len(os.Args) >= 3 && os.Args[1] == "eat" {
		pkg := os.Args[2]
		fmt.Printf("Eating %s...\n\n", pkg)
		if err := EatGo(pkg); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\nDone. Run 'go run main.go' to start server.")
		return
	}

	startServer()
}