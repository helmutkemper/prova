package importFlightStretch

import (
	"log"
	"os"
)

// Carrega o arquivo CSV e adiciona a fonte de dados, após verificação.
//   Atenção: Esta função não tem a responsabilidade de verificar dados repetidos.
func (el *CSV) LoadFileAndAppendToDataSource(filePath string) (err error) {
	el.mutex.Lock()
	defer el.mutex.Unlock()

	var filePointer *os.File

	el.clearErrorList()
	filePointer, err = el.loadFile(filePath)
	if err != nil {
		log.Printf("importFlightStretch.LoadFileAndAppendToDataSource(%v).error: %v", filePath, err)
		return
	}

	err = el.verifyFromFile(filePointer)
	if err != nil {
		log.Printf("importFlightStretch.LoadFileAndAppendToDataSource(%v).error: %v", filePath, err)
		return
	}

	filePointer, err = el.loadFile(filePath)
	if err != nil {
		log.Printf("importFlightStretch.LoadFileAndAppendToDataSource(%v).error: %v", filePath, err)
		return
	}

	err = el.fileToDataSourceAppend(filePointer)
	if err != nil {
		log.Printf("importFlightStretch.LoadFileAndAppendToDataSource(%v).error: %v", filePath, err)
		return
	}

	return
}
