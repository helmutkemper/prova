package flightSuggestion

import "commonData"

func (el *Route) GetPrice() commonData.Price {
	return el.price
}
