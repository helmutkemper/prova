package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_getLastDestinationInThisRoute(t *testing.T) {
	var destination string
	var found bool

	rlObj := Route{}
	destination, found = rlObj.getLastDestinationInThisRoute()
	if destination != "" {
		t.Fail()
		panic(nil)
	}
	if found != false {
		t.Fail()
		panic(nil)
	}

	rlObj = Route{
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
	destination, found = rlObj.getLastDestinationInThisRoute()
	if destination != "bad" {
		t.Fail()
		panic(nil)
	}
	if found != true {
		t.Fail()
		panic(nil)
	}

}
