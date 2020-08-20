package flightSuggestion

import (
	"commonData"
	"testing"
)

func TestRoute_testRulesGapsInRoute(t *testing.T) {
	var foundAGap bool

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
	foundAGap = rlObj.testRulesGapsInRoute("bad")
	if foundAGap != false {
		t.Fail()
		panic(nil)
	}
	foundAGap = rlObj.testRulesGapsInRoute("Xad")
	if foundAGap != true {
		t.Fail()
		panic(nil)
	}
}
