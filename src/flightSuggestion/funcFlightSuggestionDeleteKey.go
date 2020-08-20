package flightSuggestion

func (el *FlightSuggestion) deleteKey(deleteKey int) {
	if len(el.list) == 0 {
		return
	}

	el.list = append(el.list[:deleteKey], el.list[deleteKey+1:]...)
}
