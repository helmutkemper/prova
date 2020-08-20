package flightSuggestion

import "commonData"

// devolve os trechos de uma rota
func (el *Route) GetRoute() (route []commonData.DataFlightStretch) {
	if len(el.route) == 0 {
		el.route = make([]commonData.DataFlightStretch, 0)
	}

	return el.route
}
