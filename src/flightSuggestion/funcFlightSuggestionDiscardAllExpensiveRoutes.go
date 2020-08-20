package flightSuggestion

func (el *FlightSuggestion) discardAllExpensiveRoutes() {
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

	el.deleteKeyByList(deleteList)
}
