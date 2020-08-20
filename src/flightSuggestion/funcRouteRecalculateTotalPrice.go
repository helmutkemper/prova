package flightSuggestion

import "commonData"

func (el *Route) recalculateTotalPrice() {
	if len(el.route) == 0 {
		el.route = make([]commonData.DataFlightStretch, 0)
	}

	var newPrice commonData.Price
	for _, flightStretch := range el.route {
		newPrice += flightStretch.Price
	}

	el.price = newPrice

	return
}
