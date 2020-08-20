package flightSuggestion

// Retorna todas as rotas
func (el *FlightSuggestion) getList() (list []Route) {
	if len(el.list) == 0 {
		el.list = make([]Route, 0)
	}

	return el.list
}
