package gen

import (
	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util"
	"github.com/wangliujing/goctl/util/pathx"
)

func genCountUnscoped(table Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, countUnscopedTemplateFile, template.CountUnscoped)
	if err != nil {
		return "", "", err
	}

	output, err := util.With("countUnscoped").
		Parse(text).
		Execute(map[string]any{
			"upperStartCamelObject": camel,
		})
	if err != nil {
		return "", "", err
	}

	text, err = pathx.LoadTemplate(category, countUnscopedMethodTemplateFile, template.CountUnscopedMethod)
	if err != nil {
		return "", "", err
	}

	pageMethod, err := util.With("countUnscopedMethod").
		Parse(text).
		Execute(map[string]any{})
	if err != nil {
		return "", "", err
	}

	return output.String(), pageMethod.String(), nil
}
