package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestFlightSuggestion_deleteKey(t *testing.T) {
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
					Price:       1,
				},
			},
			price: 2,
		},
		{
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
			},
			price: 2,
		},
	}
	flObj.deleteKey(0)

	if flObj.list[0].route[0].Origin != "baa" ||
		flObj.list[0].route[0].Destination != "bab" ||
		flObj.list[0].route[0].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObj.list[0].route[1].Origin != "bab" ||
		flObj.list[0].route[1].Destination != "bac" ||
		flObj.list[0].route[1].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObj.list[0].price != commonData.Price(2) {
		t.Fail()
		panic(nil)
	}
}
