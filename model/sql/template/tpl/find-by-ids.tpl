func (m *default{{.upperStartCamelObject}}Mapper) FindByIds(ids ...{{.dataType}}) ([]model.{{.upperStartCamelObject}}PO, error) {
	result := make([]model.{{.upperStartCamelObject}}PO, 0)
	if err := m.conn.Find(&result, ids).Error; err != nil {
		return nil, err
	}
	return result, nil;
}