package flightSuggestion

// verifica se a rota tem a mesma origem e destino dos pontos solicitados
//   origin:        aeroporto de origem no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   destination:   aeroporto de destino no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   yes:           bol true para rota completa
func (el *Route) isThisRouteComplete(origin, destination string) (yes bool) {
	var length = el.getRouteLength() - 1
	if el.route[0].Origin == origin && el.route[length].Destination == destination {
		yes = true
	}

	return
}
