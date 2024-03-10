func (m *default{{.upperStartCamelObject}}Mapper) CountUnscoped() (int64, error){
	var total int64
	if err := m.conn.Unscoped().Model(&model.{{.upperStartCamelObject}}PO{}).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}