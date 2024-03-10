func (m *default{{.upperStartCamelObject}}Mapper) BatchSave(models []model.{{.upperStartCamelObject}}PO, batchSize int) (error) {
	if err := m.conn.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&models, batchSize).Error; err != nil {
		return err
	}
	return nil
}