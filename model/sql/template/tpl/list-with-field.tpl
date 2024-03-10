func (m *default{{.upperStartCamelObject}}Mapper) ListWithField(fields[]string, size int64) ([]model.{{.upperStartCamelObject}}PO, error) {
	result := make([]model.{{.upperStartCamelObject}}PO, 0)
	if err := m.conn.Select(fields).Limit(int(size)).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}