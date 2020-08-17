package testing

import (
	"errors"
	"importFlightStretch"
	"testDataSource"
	"testing"
)

func TestProcessErrorFile5(t *testing.T) {

	var err error
	var errorList []error
	var ds = testDataSource.TestDataSource{}

	var c importFlightStretch.CSV
	c.SetFieldSeparator(",")
	c.SetDataSource(&ds)

	err = c.ClearDataSourceAndLoadFile("./input-file.error.5.csv")
	if err == nil {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	errorList = c.GetErrorList()
	if len(errorList) != 4 {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	if errorList[0].Error() != "csv processing file error on line 3" {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	if errorList[1].Error() != "the IATA code of the airport of origin must be three letters" {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	if errorList[2].Error() != "the IATA code of the destination airport must be three letters" {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}

	if errorList[3].Error() != "price must be a positive number" {
		t.Fail()
		err = errors.New("verification file function has a bug")
		panic(err)
	}
}
