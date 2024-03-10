package gen

import (
	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util"
	"github.com/wangliujing/goctl/util/pathx"
	"github.com/wangliujing/goctl/util/stringx"
)

func genDeleteByIds(table Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, deleteByIdsTemplateFile, template.DeleteByIds)
	if err != nil {
		return "", "", err
	}

	output, err := util.With("deleteByIds").
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

	text, err = pathx.LoadTemplate(category, deleteByIdsMethodTemplateFile, template.DeleteByIdsMethod)
	if err != nil {
		return "", "", err
	}

	deleteByIdsMethod, err := util.With("deleteByIdsMethod").
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

	return output.String(), deleteByIdsMethod.String(), nil
}
