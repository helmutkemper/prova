package flightSuggestion

func (el *Route) isThisRouteComplete(origin, destination string) (yes bool) {
	var length = el.getRouteLength() - 1
	if el.route[0].Origin == origin && el.route[length].Destination == destination {
		yes = true
	}

	return
}
