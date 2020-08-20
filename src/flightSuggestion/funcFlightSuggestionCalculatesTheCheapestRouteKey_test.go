package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestFlightSuggestion_calculatesTheCheapestRouteKey(t *testing.T) {
	var err error
	var key int
	flObj := FlightSuggestion{}
	key, err = flObj.calculatesTheCheapestRouteKey()
	if err == nil {
		t.Fail()
		panic(nil)
	}
	if key != 0 {
		t.Fail()
		panic(nil)
	}

	flObj.list = []Route{
		{
			route: []commonData.DataFlightStretch{
				{
					Origin:      "baa",
					Destination: "bab",
					Price:       3,
				},
				{
					Origin:      "bab",
					Destination: "bac",
					Price:       1,
				},
			},
			price: 4,
		},
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
	}
	key, err = flObj.calculatesTheCheapestRouteKey()
	if err != nil {
		t.Fail()
		panic(nil)
	}
	if key != 1 {
		t.Fail()
		panic(nil)
	}

}
