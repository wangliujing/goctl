func (m *default{{.upperStartCamelObject}}Mapper) DeleteById({{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	return m.conn.Delete(&model.{{.upperStartCamelObject}}PO{
	        {{.upperStartCamelPrimaryKey}}: {{.lowerStartCamelPrimaryKey}},
	    }).Error
}