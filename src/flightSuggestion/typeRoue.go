package flightSuggestion

import (
	"commonData"
)

// obj usado para arquivar as rotas usadas por FlightSuggestion.
type Route struct {
	route []commonData.DataFlightStretch
	price commonData.Price
}
