package flightSuggestion

// Retorna a quantidade de listas de rotas
func (el *FlightSuggestion) getLength() (length int) {
	return len(el.list)
}
