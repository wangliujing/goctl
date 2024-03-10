package gen

import (
	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util"
	"github.com/wangliujing/goctl/util/pathx"
)

func genImports(table Table) (string, error) {
	text, err := pathx.LoadTemplate(category, importsTemplateFile, template.Imports)
	if err != nil {
		return "", err
	}
	var containsDeletedAt = false
	for _, field := range table.Fields {
		if field.Name.Source() == "deleted_at" {
			containsDeletedAt = true
		}
	}
	buffer, err := util.With("import").Parse(text).Execute(map[string]any{
		"containsDeletedAt": containsDeletedAt,
		//"time":              timeImport,
		"containsPQ": table.ContainsPQ,
		"data":       table,
	})
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
