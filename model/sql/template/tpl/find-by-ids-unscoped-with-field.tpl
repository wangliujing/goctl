func (m *default{{.upperStartCamelObject}}Mapper) FindByIdsUnscopedWithField(fields[]string, ids ...{{.dataType}}) ([]model.{{.upperStartCamelObject}}PO, error) {
	result := make([]model.{{.upperStartCamelObject}}PO, 0)
	if err := m.conn.Unscoped().Select(fields).Find(&result, ids).Error; err != nil {
		return nil, err
	}
	return result, nil;
}