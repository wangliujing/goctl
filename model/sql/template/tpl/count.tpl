func (m *default{{.upperStartCamelObject}}Mapper) Count() (int64, error){
	var total int64
	if err := m.conn.Model(&model.{{.upperStartCamelObject}}PO{}).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}