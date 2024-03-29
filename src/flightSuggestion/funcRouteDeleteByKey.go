package flightSuggestion

import "errors"

//todo: recalcular valores
//todo: verificar chave negativa
func (el *Route) deleteByKey(keyToDelete int) (err error) {
	if len(el.route) == 0 {
		err = errors.New("Route is empty")
		return
	}

	if keyToDelete < 0 {
		err = errors.New("key must be a positive number")
		return
	}

	if keyToDelete >= len(el.route) {
		err = errors.New("key not found")
		return
	}

	el.route = append(el.route[:keyToDelete], el.route[keyToDelete+1:]...)
	el.recalculateTotalPrice()

	return
}
