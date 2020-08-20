package flightSuggestion

func (el *FlightSuggestion) addAnotherRouteList(routeList *FlightSuggestion) {
	if len(el.list) == 0 {
		el.list = make([]Route, 0)
	}

	el.list = append(el.list, routeList.getList()...)
}
