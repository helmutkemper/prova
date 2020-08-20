package flightSuggestion

// Calcula o valor total de todas as rotas contidas na lista e deixa apenas a mais barata
//   Cuidado: Não há regra de rota fechada
func (el *FlightSuggestion) discardAllExpensiveRoutes() (err error) {
	if len(el.list) <= 1 {
		return
	}

	var deleteList = make([]int, 0)
	var cheapestKey int

	cheapestKey, _ = el.calculatesTheCheapestRouteKey()
	for k := range el.list {
		if k != cheapestKey {
			deleteList = append(deleteList, k)
		}
	}

	err = el.deleteKeyByList(deleteList)
	return
}
