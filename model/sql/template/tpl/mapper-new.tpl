func new{{.upperStartCamelObject}}Mapper(conn orm.Connection) *default{{.upperStartCamelObject}}Mapper {
	return &default{{.upperStartCamelObject}}Mapper{
		conn: conn.(*orm.GormConnection),
	}
}
