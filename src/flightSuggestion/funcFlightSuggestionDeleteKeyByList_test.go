package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestFlightSuggestion_deleteKeyByList(t *testing.T) {
	var err error

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
	err = flObj.deleteKeyByList([]int{0, 2, -4})
	if err == nil || err.Error() != "key must be a positive number" {
		t.Fail()
		panic(err)
	}

	if len(flObj.list) != 3 {
		t.Fail()
		panic(nil)
	}

	err = flObj.deleteKeyByList([]int{0, 2, 400})
	if err == nil || err.Error() != "key not found" {
		t.Fail()
		panic(err)
	}

	if len(flObj.list) != 3 {
		t.Fail()
		panic(nil)
	}

	err = flObj.deleteKeyByList([]int{0, 2})
	if err != nil {
		t.Fail()
		panic(err)
	}

	if len(flObj.list) != 1 {
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

	if flObj.list[0].price != commonData.Price(2) {
		t.Fail()
		panic(nil)
	}
}
