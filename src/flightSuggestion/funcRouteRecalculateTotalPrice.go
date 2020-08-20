package flightSuggestion

import "commonData"

// recalcula o valor total de uma rota
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
