func (m *default{{.upperStartCamelObject}}Mapper) FindById({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*model.{{.upperStartCamelObject}}PO, error) {
	result := &model.{{.upperStartCamelObject}}PO{}
	if err := m.conn.Where("{{.lowerStartCamelPrimaryKey}}=?", {{.lowerStartCamelPrimaryKey}}).Take(result).Error; err != nil {
		if err == logger.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}