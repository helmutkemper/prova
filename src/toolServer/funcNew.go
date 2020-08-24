package toolServer

import (
	"dataSourceInterface"
	"flightSuggestion"
	"restServer"
)

func New(
	route flightSuggestion.Route,
	flightSuggestionObj *flightSuggestion.FlightSuggestion,
	dataSource dataSourceInterface.DataSourceBasicInterface,
	port int,
) {

	go func(
		route flightSuggestion.Route,
		flightSuggestionObj *flightSuggestion.FlightSuggestion,
		dataSource dataSourceInterface.DataSourceBasicInterface,
		port int,
	) {

		restServer.Server(route, flightSuggestionObj, dataSource, port)
	}(route, flightSuggestionObj, dataSource, port)
}
