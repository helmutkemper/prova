package flightSuggestion

import "reflect"

// Apaga uma rota contida na lista de rotas baseado na referência de uma rota igual, não
// sendo necessário que a referencia seja de uma rota contida na mesma lista.
//   pointer: ponteiro para a rota a ser apagada
func (el *FlightSuggestion) deleteByReference(pointer *Route) {
	if len(el.list) == 0 {
		return
	}

	for deleteKey, route := range el.list {
		if reflect.DeepEqual(route, *pointer) == true {
			el.list = append(el.list[:deleteKey], el.list[deleteKey+1:]...)
			return
		}
	}
}
