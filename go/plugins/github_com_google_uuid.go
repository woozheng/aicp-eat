package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var Funcs = [
  {
    "name": "NewDCESecurity",
    "params": [
      {
        "name": "domain",
        "type": "Domain"
      },
      {
        "name": "id",
        "type": "uint32"
      }
    ]
  },
  {
    "name": "NewDCEPerson",
    "params": []
  },
  {
    "name": "NewDCEGroup",
    "params": []
  },
  {
    "name": "Domain",
    "params": []
  },
  {
    "name": "ID",
    "params": []
  },
  {
    "name": "String",
    "params": []
  },
  {
    "name": "NewHash",
    "params": [
      {
        "name": "h",
        "type": "hash.Hash"
      },
      {
        "name": "space",
        "type": "UUID"
      },
      {
        "name": "data",
        "type": "[]byte"
      },
      {
        "name": "version",
        "type": "int"
      }
    ]
  },
  {
    "name": "NewMD5",
    "params": [
      {
        "name": "space",
        "type": "UUID"
      },
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "NewSHA1",
    "params": [
      {
        "name": "space",
        "type": "UUID"
      },
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "MarshalText",
    "params": []
  },
  {
    "name": "UnmarshalText",
    "params": [
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "MarshalBinary",
    "params": []
  },
  {
    "name": "UnmarshalBinary",
    "params": [
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "NodeInterface",
    "params": []
  },
  {
    "name": "SetNodeInterface",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "NodeID",
    "params": []
  },
  {
    "name": "SetNodeID",
    "params": [
      {
        "name": "id",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "NodeID",
    "params": []
  },
  {
    "name": "Scan",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Value",
    "params": []
  },
  {
    "name": "MarshalBinary",
    "params": []
  },
  {
    "name": "UnmarshalBinary",
    "params": [
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "MarshalText",
    "params": []
  },
  {
    "name": "UnmarshalText",
    "params": [
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "MarshalJSON",
    "params": []
  },
  {
    "name": "UnmarshalJSON",
    "params": [
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "Scan",
    "params": [
      {
        "name": "src",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Value",
    "params": []
  },
  {
    "name": "UnixTime",
    "params": []
  },
  {
    "name": "GetTime",
    "params": []
  },
  {
    "name": "ClockSequence",
    "params": []
  },
  {
    "name": "SetClockSequence",
    "params": [
      {
        "name": "seq",
        "type": "int"
      }
    ]
  },
  {
    "name": "Time",
    "params": []
  },
  {
    "name": "ClockSequence",
    "params": []
  },
  {
    "name": "Error",
    "params": []
  },
  {
    "name": "IsInvalidLengthError",
    "params": [
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "Parse",
    "params": [
      {
        "name": "s",
        "type": "string"
      }
    ]
  },
  {
    "name": "ParseBytes",
    "params": [
      {
        "name": "b",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "MustParse",
    "params": [
      {
        "name": "s",
        "type": "string"
      }
    ]
  },
  {
    "name": "FromBytes",
    "params": [
      {
        "name": "b",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "Must",
    "params": [
      {
        "name": "uuid",
        "type": "UUID"
      },
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "Validate",
    "params": [
      {
        "name": "s",
        "type": "string"
      }
    ]
  },
  {
    "name": "String",
    "params": []
  },
  {
    "name": "URN",
    "params": []
  },
  {
    "name": "Variant",
    "params": []
  },
  {
    "name": "Version",
    "params": []
  },
  {
    "name": "String",
    "params": []
  },
  {
    "name": "String",
    "params": []
  },
  {
    "name": "SetRand",
    "params": [
      {
        "name": "r",
        "type": "io.Reader"
      }
    ]
  },
  {
    "name": "EnableRandPool",
    "params": []
  },
  {
    "name": "DisableRandPool",
    "params": []
  },
  {
    "name": "Strings",
    "params": []
  },
  {
    "name": "NewUUID",
    "params": []
  },
  {
    "name": "New",
    "params": []
  },
  {
    "name": "NewString",
    "params": []
  },
  {
    "name": "NewRandom",
    "params": []
  },
  {
    "name": "NewRandomFromReader",
    "params": [
      {
        "name": "r",
        "type": "io.Reader"
      }
    ]
  },
  {
    "name": "NewV6",
    "params": []
  },
  {
    "name": "NewV7",
    "params": []
  },
  {
    "name": "NewV7FromReader",
    "params": [
      {
        "name": "r",
        "type": "io.Reader"
      }
    ]
  }
]

func GetHelp() map[string]interface{} {
	var parsed []map[string]interface{}
	json.Unmarshal([]byte(Funcs), &parsed)
	return map[string]interface{}{
		"route":     "/api/github_com_google_uuid",
		"package":   "github.com/google/uuid",
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
	fmt.Fprintf(w, `{"result":"%s() called","source":"github.com/google/uuid"}`, funcName)
}
