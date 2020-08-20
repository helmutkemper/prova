package flightSuggestion

// devolve a quantidade de trechos da rota
func (el *Route) getRouteLength() (length int) {
	return len(el.route)
}
