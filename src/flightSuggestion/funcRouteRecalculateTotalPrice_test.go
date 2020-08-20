package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_recalculateTotalPrice(t *testing.T) {
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
		price: 0,
	}

	rlObjA.recalculateTotalPrice()
	if rlObjA.price != 3 {
		t.Fail()
		panic(nil)
	}
}
