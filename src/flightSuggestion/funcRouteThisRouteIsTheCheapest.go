package flightSuggestion

import "commonData"

func (el *Route) thisRouteIsTheCheapest(price commonData.Price) (yes bool) {
	return el.price > price
}
