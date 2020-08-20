package flightSuggestion

import "commonData"

func (el *Route) copy(routePointer *Route) {
	var length = routePointer.getRouteLength()
	el.route = make([]commonData.DataFlightStretch, length, length)
	copy(el.route, routePointer.GetRoute())
	el.price = routePointer.GetPrice()

	return
}
