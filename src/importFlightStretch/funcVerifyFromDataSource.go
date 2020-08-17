package importFlightStretch

import (
	"commonData"
	"errors"
	"strconv"
)

// Procura por erros contidos na fonte de dados, antes da exportação
func (el *CSV) verifyFromDataSource() (err error) {
	var errorFoundList []error
	var dataSource *[]commonData.DataFlightStretch

	dataSource, err = el.dataSource.GetAllData()
	if err != nil {
		err = errors.New("csv processing data source error: " + err.Error())
		el.errorList = append(el.errorList, err)
		return
	}

	for lineCounter, lineDataPointer := range *dataSource {
		errorFoundList = el.verifyRules(lineDataPointer.Origin, lineDataPointer.Destination, lineDataPointer.Price)
		if len(errorFoundList) != 0 {
			err = errors.New("csv processing data source error on registry " + strconv.Itoa(lineCounter))
			errorFoundList = append([]error{err}, errorFoundList...)
			el.errorList = append(el.errorList, errorFoundList...)
		}
	}

	if err != nil {
		err = errors.New("several errors were found. please check the log for more details")
	}

	return
}
