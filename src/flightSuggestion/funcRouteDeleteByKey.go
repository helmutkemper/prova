package flightSuggestion

import "errors"

//todo: recalcular valores
func (el *Route) deleteByKey(keyToDelete int) (err error) {
	if len(el.route) == 0 {
		err = errors.New("route is empty")
		return
	}

	if keyToDelete >= len(el.route) {
		err = errors.New("key not found")
		return
	}

	el.route = append(el.route[:keyToDelete], el.route[keyToDelete+1:]...)

	return
}
