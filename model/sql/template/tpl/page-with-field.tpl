func (m *default{{.upperStartCamelObject}}Mapper) PageWithField(fields[]string, page int64, size int64) (*dto.Page[model.{{.upperStartCamelObject}}PO], error){
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return dto.NewPage[model.{{.upperStartCamelObject}}PO](nil, 0, size), nil
	}
	result := make([]model.{{.upperStartCamelObject}}PO, 0)
	if err := m.conn.Select(fields).Offset(int((page - 1) * size)).Limit(int(size)).Find(&result).
		Error; err != nil {
		return nil, err;
	}
	return dto.NewPage[model.{{.upperStartCamelObject}}PO](result, total, size), nil
}