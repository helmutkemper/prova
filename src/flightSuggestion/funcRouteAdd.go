package flightSuggestion

import "commonData"

func (el *Route) Add(origin, destination string, price commonData.Price) (err error) {
	err = el.testRules(origin, destination)
	if err != nil {
		return
	}

	el.price += price

	if len(el.route) == 0 {
		el.route = make([]commonData.DataFlightStretch, 0)
	}

	var flightStretch = commonData.DataFlightStretch{
		Origin:      origin,
		Destination: destination,
		Price:       price,
	}

	el.route = append(el.route, flightStretch)

	return
}
