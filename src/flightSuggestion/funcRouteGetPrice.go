package flightSuggestion

import "commonData"

// devolve o preço da rota
func (el *Route) GetPrice() commonData.Price {
	el.recalculateTotalPrice()
	return el.price
}
