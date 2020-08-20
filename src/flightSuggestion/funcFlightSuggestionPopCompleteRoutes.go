package flightSuggestion

func (el *FlightSuggestion) popCompleteRoutes(origin, destination string) (completeRoutesList FlightSuggestion) {

	var deleteList = make([]int, 0)
	for key, route := range el.list {
		if route.isThisRouteComplete(origin, destination) == true {
			completeRoutesList.addRoute(&route)
			deleteList = append(deleteList, key)
		}
	}

	el.deleteKeyByList(deleteList)

	return
}
