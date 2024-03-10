func (m *default{{.upperStartCamelObject}}Mapper) ListUnscoped(size int64) ([]model.{{.upperStartCamelObject}}PO, error) {
	result := make([]model.{{.upperStartCamelObject}}PO, 0)
	if err := m.conn.Unscoped().Limit(int(size)).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}