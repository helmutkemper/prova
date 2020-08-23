package flightSuggestion

import (
	"commonData"
)

// obj usado para arquivar as rotas usadas por FlightSuggestion.
type Route struct {
	route []commonData.DataFlightStretch
	price commonData.Price
}

func (el Route) String() (route string) {

	for k, v := range el.route {
		if k == 0 {
			route += v.Origin
		}

		route += " - " + v.Destination
	}

	route += " > " + el.price.String()

	return
}

func (el Route) GetFlightStretchRouteAsString() (route string) {

	for k, v := range el.route {
		if k == 0 {
			route += v.Origin
		}

		route += " - " + v.Destination
	}

	return
}

func (el Route) GetFlightStretchPriceAsString() (price string) {
	return el.price.String()
}
