package flightSuggestion

// testa se a nova origem já existe na rota
func (el *Route) testRulesRoutInfiniteLoop(destination string) (infiniteLoop bool) {
	for _, flightStretch := range el.route {
		if flightStretch.Origin == destination {
			infiniteLoop = true
			return
		}
	}

	return
}
