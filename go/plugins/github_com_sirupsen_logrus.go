package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var Funcs = [
  {
    "name": "Exit",
    "params": [
      {
        "name": "code",
        "type": "int"
      }
    ]
  },
  {
    "name": "RegisterExitHandler",
    "params": [
      {
        "name": "handler",
        "type": "func()"
      }
    ]
  },
  {
    "name": "DeferExitHandler",
    "params": [
      {
        "name": "handler",
        "type": "func()"
      }
    ]
  },
  {
    "name": "Put",
    "params": [
      {
        "name": "buf",
        "type": "*bytes.Buffer"
      }
    ]
  },
  {
    "name": "Get",
    "params": []
  },
  {
    "name": "SetBufferPool",
    "params": [
      {
        "name": "bp",
        "type": "BufferPool"
      }
    ]
  },
  {
    "name": "NewEntry",
    "params": [
      {
        "name": "logger",
        "type": "*Logger"
      }
    ]
  },
  {
    "name": "Dup",
    "params": []
  },
  {
    "name": "Bytes",
    "params": []
  },
  {
    "name": "String",
    "params": []
  },
  {
    "name": "WithError",
    "params": [
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "WithContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "WithField",
    "params": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "WithFields",
    "params": [
      {
        "name": "fields",
        "type": "Fields"
      }
    ]
  },
  {
    "name": "WithTime",
    "params": [
      {
        "name": "t",
        "type": "time.Time"
      }
    ]
  },
  {
    "name": "HasCaller",
    "params": []
  },
  {
    "name": "Log",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Trace",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debug",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Print",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Info",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warn",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warning",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Error",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatal",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panic",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Logf",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Tracef",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debugf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Infof",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Printf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warnf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warningf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Errorf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatalf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panicf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Logln",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Traceln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debugln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Infoln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Println",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warnln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warningln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Errorln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatalln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panicln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "StandardLogger",
    "params": []
  },
  {
    "name": "SetOutput",
    "params": [
      {
        "name": "out",
        "type": "io.Writer"
      }
    ]
  },
  {
    "name": "SetFormatter",
    "params": [
      {
        "name": "formatter",
        "type": "Formatter"
      }
    ]
  },
  {
    "name": "SetReportCaller",
    "params": [
      {
        "name": "include",
        "type": "bool"
      }
    ]
  },
  {
    "name": "SetLevel",
    "params": [
      {
        "name": "level",
        "type": "Level"
      }
    ]
  },
  {
    "name": "GetLevel",
    "params": []
  },
  {
    "name": "IsLevelEnabled",
    "params": [
      {
        "name": "level",
        "type": "Level"
      }
    ]
  },
  {
    "name": "AddHook",
    "params": [
      {
        "name": "hook",
        "type": "Hook"
      }
    ]
  },
  {
    "name": "WithError",
    "params": [
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "WithContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "WithField",
    "params": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "WithFields",
    "params": [
      {
        "name": "fields",
        "type": "Fields"
      }
    ]
  },
  {
    "name": "WithTime",
    "params": [
      {
        "name": "t",
        "type": "time.Time"
      }
    ]
  },
  {
    "name": "Trace",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debug",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Print",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Info",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warn",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warning",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Error",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panic",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatal",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "TraceFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "DebugFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "PrintFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "InfoFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "WarnFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "WarningFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "ErrorFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "PanicFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "FatalFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "Tracef",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debugf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Printf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Infof",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warnf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warningf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Errorf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panicf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatalf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Traceln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debugln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Println",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Infoln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warnln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warningln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Errorln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panicln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatalln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Add",
    "params": [
      {
        "name": "hook",
        "type": "Hook"
      }
    ]
  },
  {
    "name": "Fire",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "entry",
        "type": "*Entry"
      }
    ]
  },
  {
    "name": "Format",
    "params": [
      {
        "name": "entry",
        "type": "*Entry"
      }
    ]
  },
  {
    "name": "Lock",
    "params": []
  },
  {
    "name": "Unlock",
    "params": []
  },
  {
    "name": "Disable",
    "params": []
  },
  {
    "name": "New",
    "params": []
  },
  {
    "name": "WithField",
    "params": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "WithFields",
    "params": [
      {
        "name": "fields",
        "type": "Fields"
      }
    ]
  },
  {
    "name": "WithError",
    "params": [
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "WithContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "WithTime",
    "params": [
      {
        "name": "t",
        "type": "time.Time"
      }
    ]
  },
  {
    "name": "Logf",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Tracef",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debugf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Infof",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Printf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warnf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warningf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Errorf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatalf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panicf",
    "params": [
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Log",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "LogFn",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "Trace",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debug",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Info",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Print",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warn",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warning",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Error",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatal",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panic",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "TraceFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "DebugFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "InfoFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "PrintFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "WarnFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "WarningFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "ErrorFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "FatalFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "PanicFn",
    "params": [
      {
        "name": "fn",
        "type": "LogFunction"
      }
    ]
  },
  {
    "name": "Logln",
    "params": [
      {
        "name": "level",
        "type": "Level"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Traceln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Debugln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Infoln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Println",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warnln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Warningln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Errorln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Fatalln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Panicln",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Exit",
    "params": [
      {
        "name": "code",
        "type": "int"
      }
    ]
  },
  {
    "name": "SetNoLock",
    "params": []
  },
  {
    "name": "SetLevel",
    "params": [
      {
        "name": "level",
        "type": "Level"
      }
    ]
  },
  {
    "name": "GetLevel",
    "params": []
  },
  {
    "name": "AddHook",
    "params": [
      {
        "name": "hook",
        "type": "Hook"
      }
    ]
  },
  {
    "name": "IsLevelEnabled",
    "params": [
      {
        "name": "level",
        "type": "Level"
      }
    ]
  },
  {
    "name": "SetFormatter",
    "params": [
      {
        "name": "formatter",
        "type": "Formatter"
      }
    ]
  },
  {
    "name": "SetOutput",
    "params": [
      {
        "name": "output",
        "type": "io.Writer"
      }
    ]
  },
  {
    "name": "SetReportCaller",
    "params": [
      {
        "name": "reportCaller",
        "type": "bool"
      }
    ]
  },
  {
    "name": "ReplaceHooks",
    "params": [
      {
        "name": "hooks",
        "type": "LevelHooks"
      }
    ]
  },
  {
    "name": "SetBufferPool",
    "params": [
      {
        "name": "pool",
        "type": "BufferPool"
      }
    ]
  },
  {
    "name": "String",
    "params": []
  },
  {
    "name": "ParseLevel",
    "params": [
      {
        "name": "lvl",
        "type": "string"
      }
    ]
  },
  {
    "name": "UnmarshalText",
    "params": [
      {
        "name": "text",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "MarshalText",
    "params": []
  },
  {
    "name": "Format",
    "params": [
      {
        "name": "entry",
        "type": "*Entry"
      }
    ]
  },
  {
    "name": "Writer",
    "params": []
  },
  {
    "name": "WriterLevel",
    "params": [
      {
        "name": "level",
        "type": "Level"
      }
    ]
  },
  {
    "name": "Writer",
    "params": []
  },
  {
    "name": "WriterLevel",
    "params": [
      {
        "name": "level",
        "type": "Level"
      }
    ]
  }
]

func GetHelp() map[string]interface{} {
	var parsed []map[string]interface{}
	json.Unmarshal([]byte(Funcs), &parsed)
	return map[string]interface{}{
		"route":     "/api/github_com_sirupsen_logrus",
		"package":   "github.com/sirupsen/logrus",
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
	fmt.Fprintf(w, `{"result":"%s() called","source":"github.com/sirupsen/logrus"}`, funcName)
}
