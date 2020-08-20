package flightSuggestion

// Faz o merge entre duas listas de rotas de voo
//   routeList: ponteiro da lista de rotas a ser copiada
func (el *FlightSuggestion) addAnotherRouteList(routeList *FlightSuggestion) {
	if len(el.list) == 0 {
		el.list = make([]Route, 0)
	}

	el.list = append(el.list, routeList.getList()...)
}
