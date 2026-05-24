package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Funcs = [
  {
    "name": "BasicAuthForRealm",
    "params": [
      {
        "name": "accounts",
        "type": "Accounts"
      },
      {
        "name": "realm",
        "type": "string"
      }
    ]
  },
  {
    "name": "BasicAuth",
    "params": [
      {
        "name": "accounts",
        "type": "Accounts"
      }
    ]
  },
  {
    "name": "BasicAuthForProxy",
    "params": [
      {
        "name": "accounts",
        "type": "Accounts"
      },
      {
        "name": "realm",
        "type": "string"
      }
    ]
  },
  {
    "name": "Copy",
    "params": []
  },
  {
    "name": "HandlerName",
    "params": []
  },
  {
    "name": "HandlerNames",
    "params": []
  },
  {
    "name": "Handler",
    "params": []
  },
  {
    "name": "FullPath",
    "params": []
  },
  {
    "name": "Next",
    "params": []
  },
  {
    "name": "IsAborted",
    "params": []
  },
  {
    "name": "Abort",
    "params": []
  },
  {
    "name": "AbortWithStatus",
    "params": [
      {
        "name": "code",
        "type": "int"
      }
    ]
  },
  {
    "name": "AbortWithStatusPureJSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "jsonObj",
        "type": "any"
      }
    ]
  },
  {
    "name": "AbortWithStatusJSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "jsonObj",
        "type": "any"
      }
    ]
  },
  {
    "name": "AbortWithError",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "Error",
    "params": [
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "Set",
    "params": [
      {
        "name": "key",
        "type": "any"
      },
      {
        "name": "value",
        "type": "any"
      }
    ]
  },
  {
    "name": "Get",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "MustGet",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetString",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetBool",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt8",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt16",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt32",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt64",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint8",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint16",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint32",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint64",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetFloat32",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetFloat64",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetTime",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetDuration",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetError",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetIntSlice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt8Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt16Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt32Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetInt64Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUintSlice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint8Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint16Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint32Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetUint64Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetFloat32Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetFloat64Slice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetStringSlice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetErrorSlice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetStringMap",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetStringMapString",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "GetStringMapStringSlice",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "Delete",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "Param",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "AddParam",
    "params": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "value",
        "type": "string"
      }
    ]
  },
  {
    "name": "Query",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "DefaultQuery",
    "params": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "defaultValue",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetQuery",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "QueryArray",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetQueryArray",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "QueryMap",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetQueryMap",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "PostForm",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "DefaultPostForm",
    "params": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "defaultValue",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetPostForm",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "PostFormArray",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetPostFormArray",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "PostFormMap",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetPostFormMap",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "FormFile",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "MultipartForm",
    "params": []
  },
  {
    "name": "SaveUploadedFile",
    "params": [
      {
        "name": "file",
        "type": "*multipart.FileHeader"
      },
      {
        "name": "dst",
        "type": "string"
      },
      {
        "name": "perm",
        "type": "...fs.FileMode"
      }
    ]
  },
  {
    "name": "Bind",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindJSON",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindXML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindQuery",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindYAML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindTOML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindPlain",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindHeader",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BindUri",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "MustBindWith",
    "params": [
      {
        "name": "obj",
        "type": "any"
      },
      {
        "name": "b",
        "type": "binding.Binding"
      }
    ]
  },
  {
    "name": "ShouldBind",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindJSON",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindXML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindQuery",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindYAML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindTOML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindPlain",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindHeader",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindUri",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindWith",
    "params": [
      {
        "name": "obj",
        "type": "any"
      },
      {
        "name": "b",
        "type": "binding.Binding"
      }
    ]
  },
  {
    "name": "ShouldBindBodyWith",
    "params": [
      {
        "name": "obj",
        "type": "any"
      },
      {
        "name": "bb",
        "type": "binding.BindingBody"
      }
    ]
  },
  {
    "name": "ShouldBindBodyWithJSON",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindBodyWithXML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindBodyWithYAML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindBodyWithTOML",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ShouldBindBodyWithPlain",
    "params": [
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ClientIP",
    "params": []
  },
  {
    "name": "RemoteIP",
    "params": []
  },
  {
    "name": "ContentType",
    "params": []
  },
  {
    "name": "IsWebsocket",
    "params": []
  },
  {
    "name": "Status",
    "params": [
      {
        "name": "code",
        "type": "int"
      }
    ]
  },
  {
    "name": "Header",
    "params": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "value",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetHeader",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "GetRawData",
    "params": []
  },
  {
    "name": "SetSameSite",
    "params": [
      {
        "name": "samesite",
        "type": "http.SameSite"
      }
    ]
  },
  {
    "name": "SetCookie",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "value",
        "type": "string"
      },
      {
        "name": "maxAge",
        "type": "int"
      },
      {
        "name": "path",
        "type": "string"
      },
      {
        "name": "domain",
        "type": "string"
      },
      {
        "name": "secure",
        "type": "bool"
      },
      {
        "name": "httpOnly",
        "type": "bool"
      }
    ]
  },
  {
    "name": "SetCookieData",
    "params": [
      {
        "name": "cookie",
        "type": "*http.Cookie"
      }
    ]
  },
  {
    "name": "Cookie",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Render",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "r",
        "type": "render.Render"
      }
    ]
  },
  {
    "name": "HTML",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "IndentedJSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "SecureJSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "JSONP",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "JSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "AsciiJSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "PureJSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "XML",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "YAML",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "TOML",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "ProtoBuf",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "BSON",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "obj",
        "type": "any"
      }
    ]
  },
  {
    "name": "String",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "format",
        "type": "string"
      },
      {
        "name": "values",
        "type": "...any"
      }
    ]
  },
  {
    "name": "Redirect",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "location",
        "type": "string"
      }
    ]
  },
  {
    "name": "Data",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "contentType",
        "type": "string"
      },
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "DataFromReader",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "contentLength",
        "type": "int64"
      },
      {
        "name": "contentType",
        "type": "string"
      },
      {
        "name": "reader",
        "type": "io.Reader"
      },
      {
        "name": "extraHeaders",
        "type": "map[string]string"
      }
    ]
  },
  {
    "name": "File",
    "params": [
      {
        "name": "filepath",
        "type": "string"
      }
    ]
  },
  {
    "name": "FileFromFS",
    "params": [
      {
        "name": "filepath",
        "type": "string"
      },
      {
        "name": "fs",
        "type": "http.FileSystem"
      }
    ]
  },
  {
    "name": "FileAttachment",
    "params": [
      {
        "name": "filepath",
        "type": "string"
      },
      {
        "name": "filename",
        "type": "string"
      }
    ]
  },
  {
    "name": "SSEvent",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "message",
        "type": "any"
      }
    ]
  },
  {
    "name": "Stream",
    "params": [
      {
        "name": "step",
        "type": "func(w io.Writer) bool"
      }
    ]
  },
  {
    "name": "Negotiate",
    "params": [
      {
        "name": "code",
        "type": "int"
      },
      {
        "name": "config",
        "type": "Negotiate"
      }
    ]
  },
  {
    "name": "NegotiateFormat",
    "params": [
      {
        "name": "offered",
        "type": "...string"
      }
    ]
  },
  {
    "name": "SetAccepted",
    "params": [
      {
        "name": "formats",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Deadline",
    "params": []
  },
  {
    "name": "Done",
    "params": []
  },
  {
    "name": "Err",
    "params": []
  },
  {
    "name": "Value",
    "params": [
      {
        "name": "key",
        "type": "any"
      }
    ]
  },
  {
    "name": "IsDebugging",
    "params": []
  },
  {
    "name": "BindWith",
    "params": [
      {
        "name": "obj",
        "type": "any"
      },
      {
        "name": "b",
        "type": "binding.Binding"
      }
    ]
  },
  {
    "name": "SetType",
    "params": [
      {
        "name": "flags",
        "type": "ErrorType"
      }
    ]
  },
  {
    "name": "SetMeta",
    "params": [
      {
        "name": "data",
        "type": "any"
      }
    ]
  },
  {
    "name": "JSON",
    "params": []
  },
  {
    "name": "MarshalJSON",
    "params": []
  },
  {
    "name": "Error",
    "params": []
  },
  {
    "name": "IsType",
    "params": [
      {
        "name": "flags",
        "type": "ErrorType"
      }
    ]
  },
  {
    "name": "Unwrap",
    "params": []
  },
  {
    "name": "ByType",
    "params": [
      {
        "name": "typ",
        "type": "ErrorType"
      }
    ]
  },
  {
    "name": "Last",
    "params": []
  },
  {
    "name": "Errors",
    "params": []
  },
  {
    "name": "JSON",
    "params": []
  },
  {
    "name": "MarshalJSON",
    "params": []
  },
  {
    "name": "String",
    "params": []
  },
  {
    "name": "Open",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Readdir",
    "params": [
      {
        "name": "_",
        "type": "int"
      }
    ]
  },
  {
    "name": "Dir",
    "params": [
      {
        "name": "root",
        "type": "string"
      },
      {
        "name": "listDirectory",
        "type": "bool"
      }
    ]
  },
  {
    "name": "Last",
    "params": []
  },
  {
    "name": "New",
    "params": [
      {
        "name": "opts",
        "type": "...OptionFunc"
      }
    ]
  },
  {
    "name": "Default",
    "params": [
      {
        "name": "opts",
        "type": "...OptionFunc"
      }
    ]
  },
  {
    "name": "Handler",
    "params": []
  },
  {
    "name": "Delims",
    "params": [
      {
        "name": "left",
        "type": "string"
      },
      {
        "name": "right",
        "type": "string"
      }
    ]
  },
  {
    "name": "SecureJsonPrefix",
    "params": [
      {
        "name": "prefix",
        "type": "string"
      }
    ]
  },
  {
    "name": "LoadHTMLGlob",
    "params": [
      {
        "name": "pattern",
        "type": "string"
      }
    ]
  },
  {
    "name": "LoadHTMLFiles",
    "params": [
      {
        "name": "files",
        "type": "...string"
      }
    ]
  },
  {
    "name": "LoadHTMLFS",
    "params": [
      {
        "name": "fs",
        "type": "http.FileSystem"
      },
      {
        "name": "patterns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "SetHTMLTemplate",
    "params": [
      {
        "name": "templ",
        "type": "*template.Template"
      }
    ]
  },
  {
    "name": "SetFuncMap",
    "params": [
      {
        "name": "funcMap",
        "type": "template.FuncMap"
      }
    ]
  },
  {
    "name": "NoRoute",
    "params": [
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "NoMethod",
    "params": [
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "Use",
    "params": [
      {
        "name": "middleware",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "With",
    "params": [
      {
        "name": "opts",
        "type": "...OptionFunc"
      }
    ]
  },
  {
    "name": "Routes",
    "params": []
  },
  {
    "name": "SetTrustedProxies",
    "params": [
      {
        "name": "trustedProxies",
        "type": "[]string"
      }
    ]
  },
  {
    "name": "Run",
    "params": [
      {
        "name": "addr",
        "type": "...string"
      }
    ]
  },
  {
    "name": "RunTLS",
    "params": [
      {
        "name": "addr",
        "type": "string"
      },
      {
        "name": "certFile",
        "type": "string"
      },
      {
        "name": "keyFile",
        "type": "string"
      }
    ]
  },
  {
    "name": "RunUnix",
    "params": [
      {
        "name": "file",
        "type": "string"
      }
    ]
  },
  {
    "name": "RunFd",
    "params": [
      {
        "name": "fd",
        "type": "int"
      }
    ]
  },
  {
    "name": "RunQUIC",
    "params": [
      {
        "name": "addr",
        "type": "string"
      },
      {
        "name": "certFile",
        "type": "string"
      },
      {
        "name": "keyFile",
        "type": "string"
      }
    ]
  },
  {
    "name": "RunListener",
    "params": [
      {
        "name": "listener",
        "type": "net.Listener"
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
    "name": "HandleContext",
    "params": [
      {
        "name": "c",
        "type": "*Context"
      }
    ]
  },
  {
    "name": "StatusCodeColor",
    "params": []
  },
  {
    "name": "LatencyColor",
    "params": []
  },
  {
    "name": "MethodColor",
    "params": []
  },
  {
    "name": "ResetColor",
    "params": []
  },
  {
    "name": "IsOutputColor",
    "params": []
  },
  {
    "name": "DisableConsoleColor",
    "params": []
  },
  {
    "name": "ForceConsoleColor",
    "params": []
  },
  {
    "name": "ErrorLogger",
    "params": []
  },
  {
    "name": "ErrorLoggerT",
    "params": [
      {
        "name": "typ",
        "type": "ErrorType"
      }
    ]
  },
  {
    "name": "Logger",
    "params": []
  },
  {
    "name": "LoggerWithFormatter",
    "params": [
      {
        "name": "f",
        "type": "LogFormatter"
      }
    ]
  },
  {
    "name": "LoggerWithWriter",
    "params": [
      {
        "name": "out",
        "type": "io.Writer"
      },
      {
        "name": "notlogged",
        "type": "...string"
      }
    ]
  },
  {
    "name": "LoggerWithConfig",
    "params": [
      {
        "name": "conf",
        "type": "LoggerConfig"
      }
    ]
  },
  {
    "name": "SetMode",
    "params": [
      {
        "name": "value",
        "type": "string"
      }
    ]
  },
  {
    "name": "DisableBindValidation",
    "params": []
  },
  {
    "name": "EnableJsonDecoderUseNumber",
    "params": []
  },
  {
    "name": "EnableJsonDecoderDisallowUnknownFields",
    "params": []
  },
  {
    "name": "Mode",
    "params": []
  },
  {
    "name": "Recovery",
    "params": []
  },
  {
    "name": "CustomRecovery",
    "params": [
      {
        "name": "handle",
        "type": "RecoveryFunc"
      }
    ]
  },
  {
    "name": "RecoveryWithWriter",
    "params": [
      {
        "name": "out",
        "type": "io.Writer"
      },
      {
        "name": "recovery",
        "type": "...RecoveryFunc"
      }
    ]
  },
  {
    "name": "CustomRecoveryWithWriter",
    "params": [
      {
        "name": "out",
        "type": "io.Writer"
      },
      {
        "name": "handle",
        "type": "RecoveryFunc"
      }
    ]
  },
  {
    "name": "Unwrap",
    "params": []
  },
  {
    "name": "WriteHeader",
    "params": [
      {
        "name": "code",
        "type": "int"
      }
    ]
  },
  {
    "name": "WriteHeaderNow",
    "params": []
  },
  {
    "name": "Write",
    "params": [
      {
        "name": "data",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "WriteString",
    "params": [
      {
        "name": "s",
        "type": "string"
      }
    ]
  },
  {
    "name": "Status",
    "params": []
  },
  {
    "name": "Size",
    "params": []
  },
  {
    "name": "Written",
    "params": []
  },
  {
    "name": "Hijack",
    "params": []
  },
  {
    "name": "CloseNotify",
    "params": []
  },
  {
    "name": "Flush",
    "params": []
  },
  {
    "name": "Pusher",
    "params": []
  },
  {
    "name": "Use",
    "params": [
      {
        "name": "middleware",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "Group",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "BasePath",
    "params": []
  },
  {
    "name": "Handle",
    "params": [
      {
        "name": "httpMethod",
        "type": "string"
      },
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "POST",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "GET",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "DELETE",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "PATCH",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "PUT",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "OPTIONS",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "HEAD",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "Any",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "methods",
        "type": "[]string"
      },
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "handlers",
        "type": "...HandlerFunc"
      }
    ]
  },
  {
    "name": "StaticFile",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "filepath",
        "type": "string"
      }
    ]
  },
  {
    "name": "StaticFileFS",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "filepath",
        "type": "string"
      },
      {
        "name": "fs",
        "type": "http.FileSystem"
      }
    ]
  },
  {
    "name": "Static",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "root",
        "type": "string"
      }
    ]
  },
  {
    "name": "StaticFS",
    "params": [
      {
        "name": "relativePath",
        "type": "string"
      },
      {
        "name": "fs",
        "type": "http.FileSystem"
      }
    ]
  },
  {
    "name": "CreateTestContext",
    "params": [
      {
        "name": "w",
        "type": "http.ResponseWriter"
      }
    ]
  },
  {
    "name": "CreateTestContextOnly",
    "params": [
      {
        "name": "w",
        "type": "http.ResponseWriter"
      },
      {
        "name": "r",
        "type": "*Engine"
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
    "name": "ByName",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Bind",
    "params": [
      {
        "name": "val",
        "type": "any"
      }
    ]
  },
  {
    "name": "WrapF",
    "params": [
      {
        "name": "f",
        "type": "http.HandlerFunc"
      }
    ]
  },
  {
    "name": "WrapH",
    "params": [
      {
        "name": "h",
        "type": "http.Handler"
      }
    ]
  },
  {
    "name": "MarshalXML",
    "params": [
      {
        "name": "e",
        "type": "*xml.Encoder"
      },
      {
        "name": "start",
        "type": "xml.StartElement"
      }
    ]
  }
]

func GetHelp() map[string]interface{} {
	var parsed []map[string]interface{}
	json.Unmarshal([]byte(Funcs), &parsed)
	return map[string]interface{}{
		"route":     "/api/github_com_gin-gonic_gin",
		"package":   "github.com/gin-gonic/gin",
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
	fmt.Fprintf(w, `{"result":"%s() called","source":"github.com/gin-gonic/gin"}`, funcName)
}
