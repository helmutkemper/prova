package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestFlightSuggestion_copyRouteAndAddFlightStretch(t *testing.T) {
	var err error

	flObj := FlightSuggestion{}
	rlObj := Route{
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
	}
	err = flObj.copyRouteAndAddFlightStretch(&rlObj, "bac", "bad", commonData.Price(3))
	if err != nil {
		t.Fail()
		panic(nil)
	}

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

	if flObj.list[0].route[2].Origin != "bac" ||
		flObj.list[0].route[2].Destination != "bad" ||
		flObj.list[0].route[2].Price != 3 {

		t.Fail()
		panic(nil)
	}

	if flObj.list[0].price != commonData.Price(5) {
		t.Fail()
		panic(nil)
	}

	err = flObj.copyRouteAndAddFlightStretch(&rlObj, "baX", "bad", commonData.Price(3))
	if err == nil || err.Error() != "the addition of this flight stretch creates a gap" {
		t.Fail()
		panic(nil)
	}

	err = flObj.copyRouteAndAddFlightStretch(&rlObj, "bac", "bab", commonData.Price(3))
	if err == nil || err.Error() != "the addition of this flight stretch creates an infinite loop" {
		t.Fail()
		panic(nil)
	}
}
