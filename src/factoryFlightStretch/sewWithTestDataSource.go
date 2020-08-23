package factoryFlightStretch

import (
	"dataSourceInterface"
	"flightSuggestion"
	"importFlightStretch"
	"testDataSource"
)

func NewWithTestDataSource(csvDataPath string) (route *flightSuggestion.Route, flightSuggestionObj *flightSuggestion.FlightSuggestion, dataSource dataSourceInterface.DataSourceBasicInterface, err error) {
	route = &flightSuggestion.Route{}
	flightSuggestionObj = &flightSuggestion.FlightSuggestion{}

	dataSource = &testDataSource.TestDataSource{}

	var csv importFlightStretch.CSV
	csv.SetDataSource(dataSource)
	csv.SetFieldSeparator(",")

	err = csv.ClearDataSourceAndLoadFile(csvDataPath)
	if err != nil {
		return
	}

	flightSuggestionObj.SetDataSource(dataSource)
	return
}
