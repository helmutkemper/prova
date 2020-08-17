package importFlightStretch

import (
	"log"
	"os"
)

// Apaga todos os dados e carrega um novo arquivo após verificação de erro.
func (el *CSV) ClearDataSourceAndLoadFile(filePath string) (err error) {
	var filePointer *os.File

	log.Printf("importFlightStretch.ClearDataSourceAndLoadFile(%v): start", filePath)

	el.clearErrorList()
	filePointer, err = el.loadFile(filePath)
	if err != nil {
		log.Printf("importFlightStretch.ClearDataSourceAndLoadFile(%v).error: %v", filePath, err)
		return
	}

	err = el.verifyFromFile(filePointer)
	if err != nil {
		log.Printf("importFlightStretch.ClearDataSourceAndLoadFile(%v).error: %v", filePath, err)
		return
	}

	el.dataSource.ClearAllData()

	filePointer, err = el.loadFile(filePath)
	if err != nil {
		log.Printf("importFlightStretch.ClearDataSourceAndLoadFile(%v).error: %v", filePath, err)
		return
	}

	err = el.fileToDataSourceAppend(filePointer)
	if err != nil {
		log.Printf("importFlightStretch.ClearDataSourceAndLoadFile(%v).error: %v", filePath, err)
		return
	}

	log.Printf("importFlightStretch.ClearDataSourceAndLoadFile(%v): success", filePath)

	return
}
