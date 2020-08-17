package importFlightStretch

import "os"

// Carrega o arquivo CSV e adiciona a fonte de dados, após verificação.
//   Atenção: Esta função não tem a responsabilidade de verificar dados repetidos.
func (el *CSV) LoadFileAndAppendToDataSource(filePath string) (err error) {
	var filePointer *os.File

	el.clearErrorList()
	filePointer, err = el.loadFile(filePath)
	if err != nil {
		return
	}

	err = el.verifyFromFile(filePointer)
	if err != nil {
		return
	}

	filePointer, err = el.loadFile(filePath)
	if err != nil {
		return
	}

	err = el.fileToDataSourceAppend(filePointer)

	return
}
