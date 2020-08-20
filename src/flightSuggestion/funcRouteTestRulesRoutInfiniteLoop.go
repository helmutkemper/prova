package flightSuggestion

func (el *Route) testRulesRoutInfiniteLoop(destination string) (infiniteLoop bool) {
	for _, flightStretch := range el.route {
		if flightStretch.Origin == destination {
			infiniteLoop = true
			return
		}
	}

	return
}
