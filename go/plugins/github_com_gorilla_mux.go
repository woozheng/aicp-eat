package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Funcs = [
  {
    "name": "Middleware",
    "params": [
      {
        "name": "handler",
        "type": "http.Handler"
      }
    ]
  },
  {
    "name": "Use",
    "params": [
      {
        "name": "mwf",
        "type": "...MiddlewareFunc"
      }
    ]
  },
  {
    "name": "CORSMethodMiddleware",
    "params": [
      {
        "name": "r",
        "type": "*Router"
      }
    ]
  },
  {
    "name": "NewRouter",
    "params": []
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "req",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "ServeHTTP",
    "params": [
      {
        "name": "w",
        "type": "http.ResponseWriter"
      },
      {
        "name": "req",
        "type": "*http.Request"
      }
    ]
  },
  {
    "name": "Get",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetRoute",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "StrictSlash",
    "params": [
      {
        "name": "value",
        "type": "bool"
      }
    ]
  },
  {
    "name": "SkipClean",
    "params": [
      {
        "name": "value",
        "type": "bool"
      }
    ]
  },
  {
    "name": "UseEncodedPath",
    "params": []
  },
  {
    "name": "NewRoute",
    "params": []
  },
  {
    "name": "Name",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Handle",
    "params": [
      {
        "name": "path",
        "type": "string"
      },
      {
        "name": "handler",
        "type": "http.Handler"
      }
    ]
  },
  {
    "name": "HandleFunc",
    "params": [
      {
        "name": "path",
        "type": "string"
      },
      {
        "name": "f",
        "type": "func(http.ResponseWriter, *http.Request)"
      }
    ]
  },
  {
    "name": "Headers",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Host",
    "params": [
      {
        "name": "tpl",
        "type": "string"
      }
    ]
  },
  {
    "name": "MatcherFunc",
    "params": [
      {
        "name": "f",
        "type": "MatcherFunc"
      }
    ]
  },
  {
    "name": "Methods",
    "params": [
      {
        "name": "methods",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Path",
    "params": [
      {
        "name": "tpl",
        "type": "string"
      }
    ]
  },
  {
    "name": "PathPrefix",
    "params": [
      {
        "name": "tpl",
        "type": "string"
      }
    ]
  },
  {
    "name": "Queries",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Schemes",
    "params": [
      {
        "name": "schemes",
        "type": "...string"
      }
    ]
  },
  {
    "name": "BuildVarsFunc",
    "params": [
      {
        "name": "f",
        "type": "BuildVarsFunc"
      }
    ]
  },
  {
    "name": "Walk",
    "params": [
      {
        "name": "walkFn",
        "type": "WalkFunc"
      }
    ]
  },
  {
    "name": "Vars",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      }
    ]
  },
  {
    "name": "CurrentRoute",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      }
    ]
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "req",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "SkipClean",
    "params": []
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "req",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "GetError",
    "params": []
  },
  {
    "name": "BuildOnly",
    "params": []
  },
  {
    "name": "Handler",
    "params": [
      {
        "name": "handler",
        "type": "http.Handler"
      }
    ]
  },
  {
    "name": "HandlerFunc",
    "params": [
      {
        "name": "f",
        "type": "func(http.ResponseWriter, *http.Request)"
      }
    ]
  },
  {
    "name": "GetHandler",
    "params": []
  },
  {
    "name": "Name",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetName",
    "params": []
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "Headers",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "HeadersRegexp",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Host",
    "params": [
      {
        "name": "tpl",
        "type": "string"
      }
    ]
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "MatcherFunc",
    "params": [
      {
        "name": "f",
        "type": "MatcherFunc"
      }
    ]
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "Methods",
    "params": [
      {
        "name": "methods",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Path",
    "params": [
      {
        "name": "tpl",
        "type": "string"
      }
    ]
  },
  {
    "name": "PathPrefix",
    "params": [
      {
        "name": "tpl",
        "type": "string"
      }
    ]
  },
  {
    "name": "Queries",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      },
      {
        "name": "match",
        "type": "*RouteMatch"
      }
    ]
  },
  {
    "name": "Schemes",
    "params": [
      {
        "name": "schemes",
        "type": "...string"
      }
    ]
  },
  {
    "name": "BuildVarsFunc",
    "params": [
      {
        "name": "f",
        "type": "BuildVarsFunc"
      }
    ]
  },
  {
    "name": "Subrouter",
    "params": []
  },
  {
    "name": "URL",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "URLHost",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "URLPath",
    "params": [
      {
        "name": "pairs",
        "type": "...string"
      }
    ]
  },
  {
    "name": "GetPathTemplate",
    "params": []
  },
  {
    "name": "GetPathRegexp",
    "params": []
  },
  {
    "name": "GetQueriesRegexp",
    "params": []
  },
  {
    "name": "GetQueriesTemplates",
    "params": []
  },
  {
    "name": "GetMethods",
    "params": []
  },
  {
    "name": "GetHostTemplate",
    "params": []
  },
  {
    "name": "GetVarNames",
    "params": []
  },
  {
    "name": "SetURLVars",
    "params": [
      {
        "name": "r",
        "type": "*http.Request"
      },
      {
        "name": "val",
        "type": "map[string]string"
      }
    ]
  }
]

func GetHelp() map[string]interface{} {
	var parsed []map[string]interface{}
	json.Unmarshal([]byte(Funcs), &parsed)
	return map[string]interface{}{
		"route":     "/api/github_com_gorilla_mux",
		"package":   "github.com/gorilla/mux",
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
	fmt.Fprintf(w, `{"result":"%s() called","source":"github.com/gorilla/mux"}`, funcName)
}
