package flightSuggestion

import "commonData"

func (el *FlightSuggestion) initNewRoute(origin, destination string, price commonData.Price) {
	if len(el.list) == 0 {
		el.list = make([]Route, 0)
	}

	var route Route
	_ = route.Add(origin, destination, price)

	el.addRoute(&route)
}
