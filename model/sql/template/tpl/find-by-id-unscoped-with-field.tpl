func (m *default{{.upperStartCamelObject}}Mapper) FindByIdUnscopedWithField(fields[]string, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*model.{{.upperStartCamelObject}}PO, error) {
	result := &model.{{.upperStartCamelObject}}PO{}
	if err := m.conn.Unscoped().Select(fields).Where("{{.lowerStartCamelPrimaryKey}}=?", {{.lowerStartCamelPrimaryKey}}).Take(result).Error; err != nil {
		if err == logger.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}