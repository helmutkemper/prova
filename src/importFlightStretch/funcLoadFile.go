package importFlightStretch

import (
	"errors"
	"os"
)

// Carrega o arquivo CSV e retorna um ponteiro de arquivo
func (el *CSV) loadFile(filePath string) (filePointer *os.File, err error) {
	filePointer, err = os.Open(filePath)
	if err != nil {
		err = errors.New("csv open file error: " + err.Error())
		el.errorList = append(el.errorList, err)
	}
	return
}
