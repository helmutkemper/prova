package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_copy(t *testing.T) {
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
		},
		price: 2,
	}

	rlObjB := Route{}
	rlObjB.copy(&rlObjA)

	if rlObjB.route[0].Origin != "baa" ||
		rlObjB.route[0].Destination != "bab" ||
		rlObjB.route[0].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if rlObjB.route[1].Origin != "bab" ||
		rlObjB.route[1].Destination != "bac" ||
		rlObjB.route[1].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if rlObjB.price != commonData.Price(2) {
		t.Fail()
		panic(nil)
	}

}
