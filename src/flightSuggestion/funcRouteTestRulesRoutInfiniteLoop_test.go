package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_testRulesRoutInfiniteLoop(t *testing.T) {
	var infiniteLoop bool

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
	infiniteLoop = rlObj.testRulesRoutInfiniteLoop("bae")
	if infiniteLoop != false {
		t.Fail()
		panic(nil)
	}
	infiniteLoop = rlObj.testRulesRoutInfiniteLoop("bab")
	if infiniteLoop != true {
		t.Fail()
		panic(nil)
	}
}
