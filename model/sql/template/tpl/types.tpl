type (
	{{.lowerStartCamelObject}}Mapper interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Mapper struct {
		conn *orm.GormConnection
	}
)
