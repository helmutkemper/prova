package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_testRules(t *testing.T) {
	var err error

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
			{
				Origin:      "bac",
				Destination: "bad",
				Price:       1,
			},
		},
		price: 0,
	}
	err = rlObj.testRules("bad", "bae")
	if err != nil {
		t.Fail()
		panic(nil)
	}

	err = rlObj.testRules("bad", "bab")
	if err == nil {
		t.Fail()
		panic(nil)
	}

	err = rlObj.testRules("Xad", "bae")
	if err == nil {
		t.Fail()
		panic(nil)
	}
}
