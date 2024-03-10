func (m *default{{.upperStartCamelObject}}Mapper) Save(model *model.{{.upperStartCamelObject}}PO) (*{{.dataType}}, error) {
	if err := m.conn.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(model).Error; err != nil {
		return nil, err
	}
	id := model.{{.upperStartCamelPrimaryKey}}
	return &id, nil
}