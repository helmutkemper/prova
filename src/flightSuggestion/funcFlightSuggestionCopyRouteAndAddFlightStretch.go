package flightSuggestion

import "commonData"

// Copia uma rota existente e adiciona um novo trecho de voo ao final da rota caso não haja errors.
// Possíveis erros:
//   the addition of this flight stretch creates an infinite loop
//   the addition of this flight stretch creates a gap
func (el *FlightSuggestion) copyRouteAndAddFlightStretch(route *Route, origin, destination string, price commonData.Price) (err error) {
	var newRoute Route

	newRoute.copy(route)
	err = newRoute.Add(origin, destination, price)
	if err != nil {
		return
	}

	el.list = append(el.list, newRoute)

	return
}
