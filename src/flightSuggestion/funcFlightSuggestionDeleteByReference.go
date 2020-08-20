package flightSuggestion

import "reflect"

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
