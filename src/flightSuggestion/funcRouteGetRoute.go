package flightSuggestion

import "commonData"

func (el *Route) GetRoute() (route []commonData.DataFlightStretch) {
	if len(el.route) == 0 {
		el.route = make([]commonData.DataFlightStretch, 0)
	}

	return el.route
}
