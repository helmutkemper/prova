package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestFlightSuggestion_initNewRoute(t *testing.T) {
	flObj := FlightSuggestion{}
	flObj.initNewRoute("aaa", "aab", commonData.Price(5))
	flObj.initNewRoute("baa", "bab", commonData.Price(3))

	if len(flObj.list) != 2 {
		t.Fail()
		panic(nil)
	}

	if flObj.list[0].route[0].Origin != "aaa" ||
		flObj.list[0].route[0].Destination != "aab" ||
		flObj.list[0].route[0].Price != 5 {

		t.Fail()
		panic(nil)
	}

	if flObj.list[0].price != commonData.Price(5) {
		t.Fail()
		panic(nil)
	}

	if flObj.list[1].route[0].Origin != "baa" ||
		flObj.list[1].route[0].Destination != "bab" ||
		flObj.list[1].route[0].Price != 3 {

		t.Fail()
		panic(nil)
	}

	if flObj.list[1].price != commonData.Price(3) {
		t.Fail()
		panic(nil)
	}
}
