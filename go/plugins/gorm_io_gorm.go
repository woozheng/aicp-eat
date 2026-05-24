package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

var Funcs = [
  {
    "name": "Association",
    "params": [
      {
        "name": "column",
        "type": "string"
      }
    ]
  },
  {
    "name": "Unscoped",
    "params": []
  },
  {
    "name": "Find",
    "params": [
      {
        "name": "out",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Append",
    "params": [
      {
        "name": "values",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Replace",
    "params": [
      {
        "name": "values",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Delete",
    "params": [
      {
        "name": "values",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Clear",
    "params": []
  },
  {
    "name": "Count",
    "params": []
  },
  {
    "name": "Create",
    "params": []
  },
  {
    "name": "Query",
    "params": []
  },
  {
    "name": "Update",
    "params": []
  },
  {
    "name": "Delete",
    "params": []
  },
  {
    "name": "Row",
    "params": []
  },
  {
    "name": "Raw",
    "params": []
  },
  {
    "name": "Execute",
    "params": [
      {
        "name": "db",
        "type": "*DB"
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
    "name": "Before",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "After",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Match",
    "params": [
      {
        "name": "fc",
        "type": "func(*DB) bool"
      }
    ]
  },
  {
    "name": "Register",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "fn",
        "type": "func(*DB)"
      }
    ]
  },
  {
    "name": "Remove",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Replace",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "fn",
        "type": "func(*DB)"
      }
    ]
  },
  {
    "name": "Before",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "After",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Register",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "fn",
        "type": "func(*DB)"
      }
    ]
  },
  {
    "name": "Remove",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Replace",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "fn",
        "type": "func(*DB)"
      }
    ]
  },
  {
    "name": "Model",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Clauses",
    "params": [
      {
        "name": "conds",
        "type": "...clause.Expression"
      }
    ]
  },
  {
    "name": "Table",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Distinct",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Select",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Omit",
    "params": [
      {
        "name": "columns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "MapColumns",
    "params": [
      {
        "name": "m",
        "type": "map[string]string"
      }
    ]
  },
  {
    "name": "Where",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Not",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Or",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Joins",
    "params": [
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "InnerJoins",
    "params": [
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Group",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Having",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Order",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Limit",
    "params": [
      {
        "name": "limit",
        "type": "int"
      }
    ]
  },
  {
    "name": "Offset",
    "params": [
      {
        "name": "offset",
        "type": "int"
      }
    ]
  },
  {
    "name": "Scopes",
    "params": [
      {
        "name": "funcs",
        "type": "...func(*DB) *DB"
      }
    ]
  },
  {
    "name": "Preload",
    "params": [
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Attrs",
    "params": [
      {
        "name": "attrs",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Assign",
    "params": [
      {
        "name": "attrs",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Unscoped",
    "params": []
  },
  {
    "name": "Raw",
    "params": [
      {
        "name": "sql",
        "type": "string"
      },
      {
        "name": "values",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Create",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "CreateInBatches",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      },
      {
        "name": "batchSize",
        "type": "int"
      }
    ]
  },
  {
    "name": "Save",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "First",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Take",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Last",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Find",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "FindInBatches",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      },
      {
        "name": "batchSize",
        "type": "int"
      },
      {
        "name": "fc",
        "type": "func(tx *DB, batch int) error"
      }
    ]
  },
  {
    "name": "FirstOrInit",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "FirstOrCreate",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Update",
    "params": [
      {
        "name": "column",
        "type": "string"
      },
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Updates",
    "params": [
      {
        "name": "values",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "UpdateColumn",
    "params": [
      {
        "name": "column",
        "type": "string"
      },
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "UpdateColumns",
    "params": [
      {
        "name": "values",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Delete",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      },
      {
        "name": "conds",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Count",
    "params": [
      {
        "name": "count",
        "type": "*int64"
      }
    ]
  },
  {
    "name": "Row",
    "params": []
  },
  {
    "name": "Rows",
    "params": []
  },
  {
    "name": "Scan",
    "params": [
      {
        "name": "dest",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Pluck",
    "params": [
      {
        "name": "column",
        "type": "string"
      },
      {
        "name": "dest",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "ScanRows",
    "params": [
      {
        "name": "rows",
        "type": "*sql.Rows"
      },
      {
        "name": "dest",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Connection",
    "params": [
      {
        "name": "fc",
        "type": "func(tx *DB) error"
      }
    ]
  },
  {
    "name": "Transaction",
    "params": [
      {
        "name": "fc",
        "type": "func(tx *DB) error"
      },
      {
        "name": "opts",
        "type": "...*sql.TxOptions"
      }
    ]
  },
  {
    "name": "Begin",
    "params": [
      {
        "name": "opts",
        "type": "...*sql.TxOptions"
      }
    ]
  },
  {
    "name": "Commit",
    "params": []
  },
  {
    "name": "Rollback",
    "params": []
  },
  {
    "name": "SavePoint",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "RollbackTo",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Exec",
    "params": [
      {
        "name": "sql",
        "type": "string"
      },
      {
        "name": "values",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "ModifyStatement",
    "params": [
      {
        "name": "stmt",
        "type": "*Statement"
      }
    ]
  },
  {
    "name": "Build",
    "params": [
      {
        "name": "",
        "type": "clause.Builder"
      }
    ]
  },
  {
    "name": "WithResult",
    "params": []
  },
  {
    "name": "G",
    "params": [
      {
        "name": "db",
        "type": "*DB"
      },
      {
        "name": "opts",
        "type": "...clause.Expression"
      }
    ]
  },
  {
    "name": "Raw",
    "params": [
      {
        "name": "sql",
        "type": "string"
      },
      {
        "name": "values",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Exec",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "sql",
        "type": "string"
      },
      {
        "name": "values",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Table",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Select",
    "params": [
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Omit",
    "params": [
      {
        "name": "columns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Set",
    "params": [
      {
        "name": "assignments",
        "type": "...clause.Assigner"
      }
    ]
  },
  {
    "name": "Create",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "r",
        "type": "*T"
      }
    ]
  },
  {
    "name": "CreateInBatches",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "r",
        "type": "*[]T"
      },
      {
        "name": "batchSize",
        "type": "int"
      }
    ]
  },
  {
    "name": "Table",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Scopes",
    "params": [
      {
        "name": "scopes",
        "type": "...func(db *Statement)"
      }
    ]
  },
  {
    "name": "Where",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Not",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Or",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Limit",
    "params": [
      {
        "name": "offset",
        "type": "int"
      }
    ]
  },
  {
    "name": "Offset",
    "params": [
      {
        "name": "offset",
        "type": "int"
      }
    ]
  },
  {
    "name": "Where",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Or",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Not",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Select",
    "params": [
      {
        "name": "columns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Omit",
    "params": [
      {
        "name": "columns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Where",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Or",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Not",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Select",
    "params": [
      {
        "name": "columns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Omit",
    "params": [
      {
        "name": "columns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Limit",
    "params": [
      {
        "name": "limit",
        "type": "int"
      }
    ]
  },
  {
    "name": "Offset",
    "params": [
      {
        "name": "offset",
        "type": "int"
      }
    ]
  },
  {
    "name": "Order",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "LimitPerRecord",
    "params": [
      {
        "name": "num",
        "type": "int"
      }
    ]
  },
  {
    "name": "Joins",
    "params": [
      {
        "name": "jt",
        "type": "clause.JoinTarget"
      },
      {
        "name": "on",
        "type": "func(db JoinBuilder, joinTable clause.Table, curTable clause.Table) error"
      }
    ]
  },
  {
    "name": "Select",
    "params": [
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Omit",
    "params": [
      {
        "name": "columns",
        "type": "...string"
      }
    ]
  },
  {
    "name": "MapColumns",
    "params": [
      {
        "name": "m",
        "type": "map[string]string"
      }
    ]
  },
  {
    "name": "Set",
    "params": [
      {
        "name": "assignments",
        "type": "...clause.Assigner"
      }
    ]
  },
  {
    "name": "Distinct",
    "params": [
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Group",
    "params": [
      {
        "name": "name",
        "type": "string"
      }
    ]
  },
  {
    "name": "Having",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Order",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Preload",
    "params": [
      {
        "name": "association",
        "type": "string"
      },
      {
        "name": "query",
        "type": "func(db PreloadBuilder) error"
      }
    ]
  },
  {
    "name": "Delete",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Update",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "value",
        "type": "any"
      }
    ]
  },
  {
    "name": "Updates",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "t",
        "type": "T"
      }
    ]
  },
  {
    "name": "Count",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "column",
        "type": "string"
      }
    ]
  },
  {
    "name": "Build",
    "params": [
      {
        "name": "builder",
        "type": "clause.Builder"
      }
    ]
  },
  {
    "name": "First",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Scan",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "result",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Last",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Take",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Find",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "FindInBatches",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "batchSize",
        "type": "int"
      },
      {
        "name": "fc",
        "type": "func(data []T, batch int) error"
      }
    ]
  },
  {
    "name": "Row",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Rows",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Update",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Create",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      }
    ]
  },
  {
    "name": "Apply",
    "params": [
      {
        "name": "config",
        "type": "*Config"
      }
    ]
  },
  {
    "name": "AfterInitialize",
    "params": [
      {
        "name": "db",
        "type": "*DB"
      }
    ]
  },
  {
    "name": "Open",
    "params": [
      {
        "name": "dialector",
        "type": "Dialector"
      },
      {
        "name": "opts",
        "type": "...Option"
      }
    ]
  },
  {
    "name": "Session",
    "params": [
      {
        "name": "config",
        "type": "*Session"
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
    "name": "Debug",
    "params": []
  },
  {
    "name": "Set",
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
    "name": "Get",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "InstanceSet",
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
    "name": "InstanceGet",
    "params": [
      {
        "name": "key",
        "type": "string"
      }
    ]
  },
  {
    "name": "Callback",
    "params": []
  },
  {
    "name": "AddError",
    "params": [
      {
        "name": "err",
        "type": "error"
      }
    ]
  },
  {
    "name": "DB",
    "params": []
  },
  {
    "name": "Expr",
    "params": [
      {
        "name": "expr",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "SetupJoinTable",
    "params": [
      {
        "name": "model",
        "type": "interface{}"
      },
      {
        "name": "field",
        "type": "string"
      },
      {
        "name": "joinTable",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Use",
    "params": [
      {
        "name": "plugin",
        "type": "Plugin"
      }
    ]
  },
  {
    "name": "ToSQL",
    "params": [
      {
        "name": "queryFn",
        "type": "func(tx *DB) *DB"
      }
    ]
  },
  {
    "name": "Migrator",
    "params": []
  },
  {
    "name": "AutoMigrate",
    "params": [
      {
        "name": "dst",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "NewPreparedStmtDB",
    "params": [
      {
        "name": "connPool",
        "type": "ConnPool"
      },
      {
        "name": "maxSize",
        "type": "int"
      },
      {
        "name": "ttl",
        "type": "time.Duration"
      }
    ]
  },
  {
    "name": "GetDBConn",
    "params": []
  },
  {
    "name": "Close",
    "params": []
  },
  {
    "name": "Reset",
    "params": []
  },
  {
    "name": "BeginTx",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "opt",
        "type": "*sql.TxOptions"
      }
    ]
  },
  {
    "name": "ExecContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "QueryContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "QueryRowContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Ping",
    "params": []
  },
  {
    "name": "GetDBConn",
    "params": []
  },
  {
    "name": "Commit",
    "params": []
  },
  {
    "name": "Rollback",
    "params": []
  },
  {
    "name": "ExecContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "QueryContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "QueryRowContext",
    "params": [
      {
        "name": "ctx",
        "type": "context.Context"
      },
      {
        "name": "query",
        "type": "string"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Ping",
    "params": []
  },
  {
    "name": "Scan",
    "params": [
      {
        "name": "rows",
        "type": "Rows"
      },
      {
        "name": "db",
        "type": "*DB"
      },
      {
        "name": "mode",
        "type": "ScanMode"
      }
    ]
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
    "name": "MarshalJSON",
    "params": []
  },
  {
    "name": "UnmarshalJSON",
    "params": [
      {
        "name": "b",
        "type": "[]byte"
      }
    ]
  },
  {
    "name": "QueryClauses",
    "params": [
      {
        "name": "f",
        "type": "*schema.Field"
      }
    ]
  },
  {
    "name": "Name",
    "params": []
  },
  {
    "name": "Build",
    "params": [
      {
        "name": "",
        "type": "clause.Builder"
      }
    ]
  },
  {
    "name": "MergeClause",
    "params": [
      {
        "name": "",
        "type": "*clause.Clause"
      }
    ]
  },
  {
    "name": "ModifyStatement",
    "params": [
      {
        "name": "stmt",
        "type": "*Statement"
      }
    ]
  },
  {
    "name": "UpdateClauses",
    "params": [
      {
        "name": "f",
        "type": "*schema.Field"
      }
    ]
  },
  {
    "name": "Name",
    "params": []
  },
  {
    "name": "Build",
    "params": [
      {
        "name": "",
        "type": "clause.Builder"
      }
    ]
  },
  {
    "name": "MergeClause",
    "params": [
      {
        "name": "",
        "type": "*clause.Clause"
      }
    ]
  },
  {
    "name": "ModifyStatement",
    "params": [
      {
        "name": "stmt",
        "type": "*Statement"
      }
    ]
  },
  {
    "name": "DeleteClauses",
    "params": [
      {
        "name": "f",
        "type": "*schema.Field"
      }
    ]
  },
  {
    "name": "Name",
    "params": []
  },
  {
    "name": "Build",
    "params": [
      {
        "name": "",
        "type": "clause.Builder"
      }
    ]
  },
  {
    "name": "MergeClause",
    "params": [
      {
        "name": "",
        "type": "*clause.Clause"
      }
    ]
  },
  {
    "name": "ModifyStatement",
    "params": [
      {
        "name": "stmt",
        "type": "*Statement"
      }
    ]
  },
  {
    "name": "WriteString",
    "params": [
      {
        "name": "str",
        "type": "string"
      }
    ]
  },
  {
    "name": "WriteByte",
    "params": [
      {
        "name": "c",
        "type": "byte"
      }
    ]
  },
  {
    "name": "WriteQuoted",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "QuoteTo",
    "params": [
      {
        "name": "writer",
        "type": "clause.Writer"
      },
      {
        "name": "field",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "Quote",
    "params": [
      {
        "name": "field",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "AddVar",
    "params": [
      {
        "name": "writer",
        "type": "clause.Writer"
      },
      {
        "name": "vars",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "AddClause",
    "params": [
      {
        "name": "v",
        "type": "clause.Interface"
      }
    ]
  },
  {
    "name": "AddClauseIfNotExists",
    "params": [
      {
        "name": "v",
        "type": "clause.Interface"
      }
    ]
  },
  {
    "name": "BuildCondition",
    "params": [
      {
        "name": "query",
        "type": "interface{}"
      },
      {
        "name": "args",
        "type": "...interface{}"
      }
    ]
  },
  {
    "name": "Build",
    "params": [
      {
        "name": "clauses",
        "type": "...string"
      }
    ]
  },
  {
    "name": "Parse",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      }
    ]
  },
  {
    "name": "ParseWithSpecialTableName",
    "params": [
      {
        "name": "value",
        "type": "interface{}"
      },
      {
        "name": "specialTableName",
        "type": "string"
      }
    ]
  },
  {
    "name": "SetColumn",
    "params": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "value",
        "type": "interface{}"
      },
      {
        "name": "fromCallbacks",
        "type": "...bool"
      }
    ]
  },
  {
    "name": "Changed",
    "params": [
      {
        "name": "fields",
        "type": "...string"
      }
    ]
  },
  {
    "name": "SelectAndOmitColumns",
    "params": [
      {
        "name": "requireCreate",
        "type": "bool"
      },
      {
        "name": "requireUpdate",
        "type": "bool"
      }
    ]
  }
]

func GetHelp() map[string]interface{} {
	var parsed []map[string]interface{}
	json.Unmarshal([]byte(Funcs), &parsed)
	return map[string]interface{}{
		"route":     "/api/gorm_io_gorm",
		"package":   "gorm.io/gorm",
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
	fmt.Fprintf(w, `{"result":"%s() called","source":"gorm.io/gorm"}`, funcName)
}
