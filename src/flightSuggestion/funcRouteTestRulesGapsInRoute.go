package flightSuggestion

// teta se a nova origem gera uma falha na rota
func (el *Route) testRulesGapsInRoute(origin string) (foundAGap bool) {
	var found bool
	var lastDestination string

	lastDestination, found = el.getLastDestinationInThisRoute()
	if found == true && lastDestination != origin {
		foundAGap = true
	}

	return
}
