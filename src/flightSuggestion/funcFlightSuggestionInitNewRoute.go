package flightSuggestion

import "commonData"

// Inicializa uma rota
//   origin:        aeroporto de origem no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   destination:   aeroporto de destino no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   price:         preço conforme o tipo commonData.Price()
func (el *FlightSuggestion) initNewRoute(origin, destination string, price commonData.Price) {
	if len(el.list) == 0 {
		el.list = make([]Route, 0)
	}

	var route Route
	_ = route.Add(origin, destination, price)

	el.addRoute(&route)
}
