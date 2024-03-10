package {{.pkg}}


import "github.com/wangliujing/foundation-framework/orm"

var _ {{.upperStartCamelObject}}Mapper = (*custom{{.upperStartCamelObject}}Mapper)(nil)

type (
	{{.upperStartCamelObject}}Mapper interface {
		{{.lowerStartCamelObject}}Mapper
	}

	custom{{.upperStartCamelObject}}Mapper struct {
		*default{{.upperStartCamelObject}}Mapper
	}
)

func New{{.upperStartCamelObject}}Mapper(conn orm.Connection) {{.upperStartCamelObject}}Mapper {
	return &custom{{.upperStartCamelObject}}Mapper{
		default{{.upperStartCamelObject}}Mapper: new{{.upperStartCamelObject}}Mapper(conn),
	}
}
