package restServer

import (
	"encoding/json"
	"flightSuggestion"
)

func (el output) MarshalJSON() (data []byte, err error) {
	if el.err != nil {

		el.Error = make([]string, 1)
		el.Error[0] = el.err.Error()

		el.Object = nil
		return json.Marshal(struct {
			Error   []string
			Success bool
			Length  int
			Object  []object
			route   flightSuggestion.Route
			err     error
		}{
			Error:  el.Error,
			Object: []object{},
		})
	}

	el.Success = true
	el.Length = 1
	el.Object = make([]object, 1)
	el.Object[0].Route = el.route.GetFlightStretchRouteAsString()
	el.Object[0].Price = el.route.GetPrice()

	return json.Marshal(struct {
		Error   []string
		Success bool
		Length  int
		Object  []object
		route   flightSuggestion.Route
		err     error
	}{
		Error:   []string{},
		Success: el.Success,
		Length:  el.Length,
		Object:  el.Object,
	})
}
