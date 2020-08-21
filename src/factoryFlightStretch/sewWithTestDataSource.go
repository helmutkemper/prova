package factoryFlightStretch

import (
	"dataSourceInterface"
	"flightSuggestion"
	"importFlightStretch"
	"testDataSource"
)

func NewWithTestDataSource(csvDataPath string) (flightSuggestionObj *flightSuggestion.FlightSuggestion, err error) {
	flightSuggestionObj = &flightSuggestion.FlightSuggestion{}

	var dataSource dataSourceInterface.DataSourceBasicInterface
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
