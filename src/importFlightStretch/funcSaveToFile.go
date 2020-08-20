package importFlightStretch

import (
	"commonData"
	"errors"
	"os"
)

// salva o arquivo CSV
func (el *CSV) saveToFile(filePath string) (err error) {
	var filePointer *os.File
	var dataSource *[]commonData.DataFlightStretch

	filePointer, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		err = errors.New("csv save file error: " + err.Error())
		el.errorList = append(el.errorList, err)
		return
	}

	defer func() {
		var err error
		err = filePointer.Close()
		if err != nil {
			//todo: melhorar
			panic(err)
		}
	}()

	dataSource, err = el.dataSource.GetAllData()
	if err != nil {
		err = errors.New("csv save file error: " + err.Error())
		el.errorList = append(el.errorList, err)
		return
	}

	var length = len(*dataSource) - 1
	for k, lineData := range *dataSource {
		var line = el.dataToString(lineData)
		_, err = filePointer.WriteString(line)
		if err != nil {
			err = errors.New("csv save file error: " + err.Error())
			el.errorList = append(el.errorList, err)
			return
		}

		if k == length {
			break
		}
		_, err = filePointer.WriteString("\r\n")
		if err != nil {
			err = errors.New("csv save file error: " + err.Error())
			el.errorList = append(el.errorList, err)
			return
		}
	}

	return
}
