package importFlightStretch

import "log"

// Salva o conteúdo da fonte de dados em um arquivo CSV
func (el *CSV) SaveDataSourceToFile(filePath string) (err error) {
	el.mutex.Lock()
	defer el.mutex.Unlock()

	err = el.verifyFromDataSource()
	if err != nil {
		log.Printf("importFlightStretch.SaveDataSourceToFile(%v).error: %v", filePath, err)
		return
	}

	err = el.saveToFile(filePath)
	if err != nil {
		log.Printf("importFlightStretch.SaveDataSourceToFile(%v).error: %v", filePath, err)
		return
	}

	return
}
