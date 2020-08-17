package factoryFlightStretch

import (
	"dataSourceInterface"
	"importFlightStretch"
)

func NewWithFieldSeparator(fieldSeparator string, dataSource dataSourceInterface.DataSourceBasicInterface) (pointer *importFlightStretch.CSV) {
	pointer = &importFlightStretch.CSV{}
	pointer.SetDataSource(dataSource)
	pointer.SetFieldSeparator(fieldSeparator)

	return
}
