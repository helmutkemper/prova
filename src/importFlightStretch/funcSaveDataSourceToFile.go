package importFlightStretch

// Salva o conteúdo da fonte de dados em um arquivo CSV
func (el *CSV) SaveDataSourceToFile(filePath string) (err error) {
	err = el.verifyFromDataSource()
	if err != nil {
		return
	}

	err = el.saveToFile(filePath)

	return
}
