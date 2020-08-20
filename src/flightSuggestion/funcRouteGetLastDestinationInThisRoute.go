package flightSuggestion

// retorna o ultimo destino da rota
func (el *Route) getLastDestinationInThisRoute() (destination string, found bool) {
	if len(el.route) == 0 {
		return
	}

	found = true
	destination = el.route[len(el.route)-1].Destination

	return
}
