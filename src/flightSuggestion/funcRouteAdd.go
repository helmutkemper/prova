package flightSuggestion

import "commonData"

// adiciona um nov trecho a uma rota caso não haja erros
//   origin:        aeroporto de origem no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   destination:   aeroporto de destino no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   price:         preço do trecho conforme commonData.Price()
//   err:           erro no padrão Golang
func (el *Route) Add(origin, destination string, price commonData.Price) (err error) {
	err = el.TestRules(origin, destination)
	if err != nil {
		return
	}

	el.price += price

	if len(el.route) == 0 {
		el.route = make([]commonData.DataFlightStretch, 0)
	}

	var flightStretch = commonData.DataFlightStretch{
		Origin:      origin,
		Destination: destination,
		Price:       price,
	}

	el.route = append(el.route, flightStretch)

	return
}
