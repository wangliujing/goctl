func (m *default{{.upperStartCamelObject}}Mapper) FindByIdUnscoped({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*model.{{.upperStartCamelObject}}PO, error) {
	result := &model.{{.upperStartCamelObject}}PO{}
	if err := m.conn.Unscoped().Where("{{.lowerStartCamelPrimaryKey}}=?", {{.lowerStartCamelPrimaryKey}}).Take(result).Error; err != nil {
		if err == logger.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}