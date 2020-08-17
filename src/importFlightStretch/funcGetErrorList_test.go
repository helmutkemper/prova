package importFlightStretch

import (
	"errors"
	"fmt"
	"testDataSource"
)

func ExampleCSV_GetErrorList() {
	var err error
	var errorList []error
	var ds = testDataSource.TestDataSource{}

	var c CSV
	c.SetFieldSeparator(",")
	c.SetDataSource(&ds)

	err = c.ClearDataSourceAndLoadFile("./testing/input-file.error.1.csv")
	if err == nil {
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	errorList = c.GetErrorList()
	for _, err = range errorList {
		fmt.Printf("%v\n", err.Error())
	}

	// Output:
	// csv processing file error on line 1
	// the IATA code of the airport of origin must be three letters
}
