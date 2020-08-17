package testing

import (
  "commonData"
  "errors"
  "importFlightStretch"
  "testDataSource"
  "testing"
)

func Test(t *testing.T) {
  ProcessOK(t)
  ProcessError1(t)
  ProcessError2(t)
  ProcessError3(t)
  ProcessError4(t)
  ProcessError5(t)
}

func ProcessOK(t *testing.T) {

  var err error
  var ds = testDataSource.TestDataSource{}
  var list []commonData.DataSegment

  var c importFlightStretch.CSV
  c.SetFieldSeparator(",")
  c.SetDataSource(&ds)

  err = c.ClearDataSourceAndLoadFile("./input-file.correct.csv")
  if err != nil {
    t.Fail()
    panic(err)
  }

  list, err = ds.GetSegmentByDestinationIataCode("BRC")
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

func ProcessError1(t *testing.T) {

  var err error
  var errorList []error
  var ds = testDataSource.TestDataSource{}

  var c importFlightStretch.CSV
  c.SetFieldSeparator(",")
  c.SetDataSource(&ds)

  err = c.ClearDataSourceAndLoadFile("./input-file.error.1.csv")
  if err == nil {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }

  errorList = c.GetErrorList()
  if len(errorList) != 2 {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }

  if errorList[0].Error() != "csv processing file error on line 1" {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }

  if errorList[1].Error() != "the IATA code of the airport of origin must be three letters" {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }
}

func ProcessError2(t *testing.T) {

  var err error
  var errorList []error
  var ds = testDataSource.TestDataSource{}

  var c importFlightStretch.CSV
  c.SetFieldSeparator(",")
  c.SetDataSource(&ds)

  err = c.ClearDataSourceAndLoadFile("./input-file.error.2.csv")
  if err == nil {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }

  errorList = c.GetErrorList()
  if len(errorList) != 2 {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }

  if errorList[0].Error() != "csv processing file error on line 3" {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }

  if errorList[1].Error() != "the IATA code of the destination airport must be three letters" {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }
}

func ProcessError3(t *testing.T) {

  var err error
  var errorList []error
  var ds = testDataSource.TestDataSource{}

  var c importFlightStretch.CSV
  c.SetFieldSeparator(",")
  c.SetDataSource(&ds)

  err = c.ClearDataSourceAndLoadFile("./input-file.error.3.csv")
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

  if errorList[0].Error() != "csv processing file error on line 4 strconv.ParseInt: parsing \"REC\": invalid syntax" {
    t.Fail()
    err = errors.New("verification file function has a bug")
    panic(err)
  }
}

func ProcessError4(t *testing.T) {

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

func ProcessError5(t *testing.T) {

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

