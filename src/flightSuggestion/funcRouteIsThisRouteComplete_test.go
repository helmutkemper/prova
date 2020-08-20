package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_isThisRouteComplete(t *testing.T) {
	var routeComplete bool

	rlObjA := Route{
		route: []commonData.DataFlightStretch{
			{
				Origin:      "baa",
				Destination: "bab",
				Price:       1,
			},
			{
				Origin:      "bab",
				Destination: "bac",
				Price:       1,
			},
			{
				Origin:      "bac",
				Destination: "bad",
				Price:       1,
			},
		},
		price: 3,
	}

	routeComplete = rlObjA.isThisRouteComplete("baa", "Xad")
	if routeComplete != false {
		t.Fail()
		panic(nil)
	}

	routeComplete = rlObjA.isThisRouteComplete("baa", "bad")
	if routeComplete != true {
		t.Fail()
		panic(nil)
	}
}
