package gen

import (
	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util"
	"github.com/wangliujing/goctl/util/pathx"
	"github.com/wangliujing/goctl/util/stringx"
)

func genFindByIdsUnscopedWithField(table Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, findByIdsUnscopedWithFieldTemplateFile, template.FindByIdsUnscopedWithField)
	if err != nil {
		return "", "", err
	}

	output, err := util.With("findByIdsUnscopedWithField").
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

	text, err = pathx.LoadTemplate(category, findByIdsUnscopedWithFieldMethodTemplateFile, template.FindByIdsUnscopedWithFieldMethod)
	if err != nil {
		return "", "", err
	}

	findByIdsUnscopedWithFieldMethod, err := util.With("findByIdsUnscopedWithFieldMethod").
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

	return output.String(), findByIdsUnscopedWithFieldMethod.String(), nil
}
