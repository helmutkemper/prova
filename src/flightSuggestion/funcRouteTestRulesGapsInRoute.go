package flightSuggestion

func (el *Route) testRulesGapsInRoute(origin string) (foundAGap bool) {
	var found bool
	var lastDestination string

	lastDestination, found = el.getLastDestinationInThisRoute()
	if found == true && lastDestination != origin {
		foundAGap = true
	}

	return
}
