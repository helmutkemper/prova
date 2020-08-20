package testing

import (
	"commonData"
	"errors"
	"importFlightStretch"
	"testDataSource"
	"testing"
)

func TestProcessOK(t *testing.T) {

	var err error
	var ds = testDataSource.TestDataSource{}
	var list []commonData.DataFlightStretch

	var c importFlightStretch.CSV
	c.SetFieldSeparator(",")
	c.SetDataSource(&ds)

	err = c.ClearDataSourceAndLoadFile("./input-file.correct.csv")
	if err != nil {
		t.Fail()
		panic(err)
	}

	list, err = ds.GetFlightStretchByDestinationIataCode("BRC")
	if err != nil {
		t.Fail()
		panic(err)
	}

	if len(list) != 1 {
		t.Fail()
		panic(errors.New("data length error"))
	}

	if (list)[0].Origin != "GRU" {
		t.Fail()
		panic(errors.New("data source segment value error"))
	}

	if (list)[0].Destination != "BRC" {
		t.Fail()
		panic(errors.New("data destination segment value error"))
	}

	if (list)[0].Price != 10 {
		t.Fail()
		panic(errors.New("data price error"))
	}
}
