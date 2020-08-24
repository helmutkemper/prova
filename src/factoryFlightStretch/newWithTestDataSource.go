package factoryFlightStretch

import (
	"dataSourceInterface"
	"flightSuggestion"
	"importFlightStretch"
	"log"
	"testDataSource"
)

func NewWithTestDataSourceAndAutoCsvSave(csvDataPath string) (route *flightSuggestion.Route, flightSuggestionObj *flightSuggestion.FlightSuggestion, dataSource dataSourceInterface.DataSourceBasicInterface, err error) {
	route = &flightSuggestion.Route{}
	flightSuggestionObj = &flightSuggestion.FlightSuggestion{}

	dataSource = &testDataSource.TestDataSource{
		Event: make(chan bool, 1),
	}

	var csv importFlightStretch.CSV
	csv.SetDataSource(dataSource)
	csv.SetFieldSeparator(",")

	err = csv.ClearDataSourceAndLoadFile(csvDataPath)
	if err != nil {
		return
	}

	flightSuggestionObj.SetDataSource(dataSource)

	go func(dataSource dataSourceInterface.DataSourceBasicInterface, ch *chan bool) {
		var err error
		for {
			select {
			case <-*ch:
				err = csv.SaveDataSourceToFile(csvDataPath)
				if err != nil {
					log.Fatalf("csv.SaveDataSourceToFile().error: %v", err)
				}
			}
		}
	}(dataSource, dataSource.GetEventChannel())

	return
}
