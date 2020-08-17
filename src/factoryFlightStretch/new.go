package factoryFlightStretch

import (
	"dataSourceInterface"
	"importFlightStretch"
)

func New(dataSource dataSourceInterface.DataSourceBasicInterface) (pointer *importFlightStretch.CSV) {
	pointer = &importFlightStretch.CSV{}
	pointer.SetDataSource(dataSource)
	pointer.SetFieldSeparator(",")

	return
}
