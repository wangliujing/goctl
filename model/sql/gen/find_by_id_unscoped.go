package gen

import (
	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util"
	"github.com/wangliujing/goctl/util/pathx"
	"github.com/wangliujing/goctl/util/stringx"
)

func genFindByIdUnscoped(table Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, findByIdUnscopedTemplateFile, template.FindByIdUnscoped)
	if err != nil {
		return "", "", err
	}

	output, err := util.With("findByIdUnscoped").
		Parse(text).
		Execute(map[string]any{
			"upperStartCamelObject":     camel,
			"lowerStartCamelObject":     stringx.From(camel).Untitle(),
			"lowerStartCamelPrimaryKey": util.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"upperStartCamelPrimaryKey": table.PrimaryKey.Name.ToCamel(),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	text, err = pathx.LoadTemplate(category, findByIdUnscopedMethodTemplateFile, template.FindByIdUnscopedMethod)
	if err != nil {
		return "", "", err
	}

	findByIdUnscopedMethod, err := util.With("findByIdUnscopedMethod").
		Parse(text).
		Execute(map[string]any{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": util.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return output.String(), findByIdUnscopedMethod.String(), nil
}
