func (m *default{{.upperStartCamelObject}}Mapper) List(size int64) ([]model.{{.upperStartCamelObject}}PO, error) {
	result := make([]model.{{.upperStartCamelObject}}PO, 0)
	if size < 0 {
	    if err := m.conn.Find(&result).Error; err != nil {
    		return nil, err
    	}
	} else {
		if err := m.conn.Limit(int(size)).Find(&result).Error; err != nil {
    		return nil, err
    	}
	}
	return result, nil
}