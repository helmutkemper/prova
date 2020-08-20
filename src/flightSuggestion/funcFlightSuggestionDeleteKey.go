package flightSuggestion

import "errors"

//todo: limite das chaves
func (el *FlightSuggestion) deleteKey(keyToDelete int) (err error) {
	if len(el.list) == 0 {
		return
	}

	if keyToDelete < 0 {
		err = errors.New("key must be a positive number")
		return
	}

	if keyToDelete >= len(el.list) {
		err = errors.New("key not found")
		return
	}

	el.list = append(el.list[:keyToDelete], el.list[keyToDelete+1:]...)

	return
}
