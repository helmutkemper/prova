package importFlightStretch

import (
	"commonData"
	"fmt"
	"testDataSource"
)

func ExampleCSV_ClearDataSourceAndLoadFile() {
	var err error
	var ds = testDataSource.TestDataSource{}
	var list []commonData.DataFlightStretch

	var c CSV
	c.SetFieldSeparator(",")
	c.SetDataSource(&ds)

	err = c.ClearDataSourceAndLoadFile("./testing/input-file.correct.csv")
	if err != nil {
		panic(err)
	}

	list, err = ds.GetFlightStretchByDestinationIataCode("BRC")
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %-v", list)

	// Output:
	// data: [{GRU BRC 10}]
}
