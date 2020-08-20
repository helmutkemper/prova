package flightSuggestion

import (
	"testing"
)

func TestRoute_Add(t *testing.T) {
	var err error
	var rt = Route{}
	err = rt.Add("GRU", "BRC", 0)
	if err != nil {
		t.Fail()
	}

	err = rt.Add("BRC", "SCL", 0)
	if err != nil {
		t.Fail()
	}

	err = rt.Add("SCL", "ORL", 0)
	if err != nil {
		t.Fail()
	}

	err = rt.Add("ORL", "BRC", 0)
	if err == nil || err.Error() != "the addition of this flight stretch creates an infinite loop" {
		t.Fail()
	}

	//
	err = rt.Add("REC", "BRL", 0)
	if err == nil || err.Error() != "the addition of this flight stretch creates a gap" {
		t.Fail()
	}
}
