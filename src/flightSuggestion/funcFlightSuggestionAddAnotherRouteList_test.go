package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestFlightSuggestion_addAnotherRouteList(t *testing.T) {
	flObjA := FlightSuggestion{}
	flObjA.list = []Route{
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
	}
	flObjB := FlightSuggestion{}
	flObjB.list = []Route{
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
	flObjB.addAnotherRouteList(&flObjA)

	if flObjB.list[0].route[0].Origin != "baa" ||
		flObjB.list[0].route[0].Destination != "bab" ||
		flObjB.list[0].route[0].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObjB.list[0].route[1].Origin != "bab" ||
		flObjB.list[0].route[1].Destination != "bac" ||
		flObjB.list[0].route[1].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObjB.list[0].price != commonData.Price(2) {
		t.Fail()
		panic(nil)
	}

	if flObjB.list[1].route[0].Origin != "aaa" ||
		flObjB.list[1].route[0].Destination != "aab" ||
		flObjB.list[1].route[0].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObjB.list[1].route[1].Origin != "aab" ||
		flObjB.list[1].route[1].Destination != "aac" ||
		flObjB.list[1].route[1].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if flObjB.list[1].price != commonData.Price(2) {
		t.Fail()
		panic(nil)
	}
}
