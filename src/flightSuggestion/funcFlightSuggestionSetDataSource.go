package flightSuggestion

import "dataSourceInterface"

func (el *FlightSuggestion) SetDataSource(dataSource dataSourceInterface.DataSourceBasicInterface) {
	el.dataSource = dataSource
}
