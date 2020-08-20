package flightSuggestion

// Adiciona uma rota a lista
func (el *FlightSuggestion) addRoute(route *Route) {
	if len(el.list) == 0 {
		el.list = make([]Route, 0)
	}

	el.list = append(el.list, *route)
}
