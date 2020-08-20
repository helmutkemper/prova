package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestFlightSuggestion_discardAllExpensiveRoutes(t *testing.T) {
	flObj := FlightSuggestion{}
	flObj.list = []Route{
		{
			route: []commonData.DataFlightStretch{
				{
					Origin:      "aaa",
					Destination: "aab",
					Price:       1,
				},
				{
					Origin:      "aab",
					Destination: "aac",
					Price:       2,
				},
			},
			price: 3,
		},
		{
			route: []commonData.DataFlightStretch{
				{
					Origin:      "baa",
					Destination: "bab",
					Price:       2,
				},
				{
					Origin:      "bab",
					Destination: "bac",
					Price:       4,
				},
			},
			price: 6,
		},
		{
			route: []commonData.DataFlightStretch{
				{
					Origin:      "caa",
					Destination: "cab",
					Price:       1,
				},
				{
					Origin:      "cab",
					Destination: "cac",
					Price:       1,
				},
			},
			price: 2,
		},
	}
	flObj.discardAllExpensiveRoutes()

	if len(flObj.list) != 1 {
		t.Fail()
		panic(nil)
	}

	if flObj.list[0].route[0].Origin != "caa" ||
		flObj.list[0].route[0].Destination != "cab" ||
		flObj.list[0].route[0].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObj.list[0].route[1].Origin != "cab" ||
		flObj.list[0].route[1].Destination != "cac" ||
		flObj.list[0].route[1].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObj.list[0].price != commonData.Price(2) {
		t.Fail()
		panic(nil)
	}
}
