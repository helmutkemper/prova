package testing

import (
	"errors"
	"importFlightStretch"
	"testDataSource"
	"testing"
)

func TestProcessErrorFile4(t *testing.T) {

	var err error
	var errorList []error
	var ds = testDataSource.TestDataSource{}

	var c importFlightStretch.CSV
	c.SetFieldSeparator(",")
	c.SetDataSource(&ds)

	err = c.ClearDataSourceAndLoadFile("./input-file.error.4.csv")
	if err == nil {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	errorList = c.GetErrorList()
	if len(errorList) != 1 {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	if errorList[0].Error() != "csv processing file error on line 3 strconv.ParseInt: parsing \"\": invalid syntax" {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}
}
