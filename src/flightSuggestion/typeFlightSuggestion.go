package flightSuggestion

import "dataSourceInterface"

// Este obj sugere rotas de voos.
//
// Para mais informações, veja a função FindCheapestRoute()
type FlightSuggestion struct {
	list       []Route
	dataSource dataSourceInterface.DataSourceBasicInterface
}
