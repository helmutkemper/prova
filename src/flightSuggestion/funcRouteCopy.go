package flightSuggestion

import "commonData"

// copia todos os dados de uma rota existente
//   routePointer: ponteiro da rota a ser copiada
func (el *Route) copy(routePointer *Route) {
	var length = routePointer.getRouteLength()
	el.route = make([]commonData.DataFlightStretch, length, length)
	copy(el.route, routePointer.GetRoute())
	el.price = routePointer.GetPrice()

	return
}
