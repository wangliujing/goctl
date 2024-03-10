package model
import (
    {{if .sql}}"database/sql"{{end}}
	{{if .time}}"time"{{end}}
)

type {{.upperStartCamelObject}}PO struct {
		{{.fields}}
}

func (m *{{.upperStartCamelObject}}PO) TableName() string {
	return "{{.tableName}}"
}
