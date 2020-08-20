package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_deleteByKeyList(t *testing.T) {
	var err error

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
		price: 3,
	}

	err = rlObjA.deleteByKeyList([]int{0, 1, 100})
	if err == nil || err.Error() != "key not found" {
		t.Fail()
		panic(nil)
	}

	err = rlObjA.deleteByKeyList([]int{0, 1, -100})
	if err == nil || err.Error() != "key must be a positive number" {
		t.Fail()
		panic(nil)
	}

	err = rlObjA.deleteByKeyList([]int{1})
	if err != nil {
		t.Fail()
		panic(nil)
	}

	if rlObjA.route[0].Origin != "baa" ||
		rlObjA.route[0].Destination != "bab" ||
		rlObjA.route[0].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if rlObjA.route[1].Origin != "bac" ||
		rlObjA.route[1].Destination != "bad" ||
		rlObjA.route[1].Price != 1 {

		t.Fail()
		panic(nil)
	}

	if rlObjA.price != commonData.Price(2) {
		t.Fail()
		panic(nil)
	}

}
