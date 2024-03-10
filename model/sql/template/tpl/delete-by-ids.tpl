func (m *default{{.upperStartCamelObject}}Mapper) DeleteByIds(ids ...{{.dataType}}) error {
	return m.conn.Delete(&model.{{.upperStartCamelObject}}PO{}, ids).Error
}