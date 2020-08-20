package flightSuggestion

import "commonData"

// verifica o preço da rota
//   price: preço baseado em commonData.Price()
//   yes:   bool true se a rota for mais barata
func (el *Route) thisRouteIsTheCheapest(price commonData.Price) (yes bool) {
	return el.price > price
}
