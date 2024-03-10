package gen

import (
	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util"
	"github.com/wangliujing/goctl/util/pathx"
)

func genCount(table Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, countTemplateFile, template.Count)
	if err != nil {
		return "", "", err
	}

	output, err := util.With("count").
		Parse(text).
		Execute(map[string]any{
			"upperStartCamelObject": camel,
		})
	if err != nil {
		return "", "", err
	}

	text, err = pathx.LoadTemplate(category, countMethodTemplateFile, template.CountMethod)
	if err != nil {
		return "", "", err
	}

	pageMethod, err := util.With("countMethod").
		Parse(text).
		Execute(map[string]any{})
	if err != nil {
		return "", "", err
	}

	return output.String(), pageMethod.String(), nil
}
