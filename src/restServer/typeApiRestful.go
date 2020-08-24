package restServer

import (
	"dataSourceInterface"
	"flightSuggestion"
)

type apiRestful struct {
	flightSuggestion *flightSuggestion.FlightSuggestion
	route            flightSuggestion.Route
	dataSource       dataSourceInterface.DataSourceBasicInterface
}
