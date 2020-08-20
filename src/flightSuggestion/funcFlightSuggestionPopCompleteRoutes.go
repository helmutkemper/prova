package flightSuggestion

// Devolve um novo objeto com todas as rotas completas
//   origin:             aeroporto de origem no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   destination:        aeroporto de destino no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   completeRoutesList: novo objeto com todas as rotas de voo completas
//   err:                erro no padrão Golang
func (el *FlightSuggestion) popCompleteRoutes(origin, destination string) (completeRoutesList FlightSuggestion, err error) {

	var deleteList = make([]int, 0)
	for key, route := range el.list {
		if route.isThisRouteComplete(origin, destination) == true {
			completeRoutesList.addRoute(&route)
			deleteList = append(deleteList, key)
		}
	}

	err = el.deleteKeyByList(deleteList)
	return
}
