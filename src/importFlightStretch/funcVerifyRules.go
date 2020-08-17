package importFlightStretch

import "errors"

// view-source:https://raw.githubusercontent.com/datasets/airport-codes/master/data/airport-codes.csv
func (el *CSV) verifyRules(origin, destination string, price int) (errorList []error) {
  errorList = make([]error, 0)
  var err error

  var originEmpty      = origin == ""
  var destinationEmpty = destination == ""
  var priceEmpty       = price == 0

  if originEmpty || destinationEmpty || priceEmpty {
    err = errors.New("the file must contain three fields per line: IATA code from airport of origin with three letters; IATA code from destination airport with three letters; price of the segment in numeric format")
    errorList = append(errorList, err)
  }

  if len(origin) != 3 {
    err = errors.New("the IATA code of the airport of origin must be three letters")
    errorList = append(errorList, err)
  }

  if len(destination) != 3 {
    err = errors.New("the IATA code of the destination airport must be three letters")
    errorList = append(errorList, err)
  }

  if price < 0 {
    err = errors.New("price must be a positive number")
    errorList = append(errorList, err)
  }

  if err != nil {
    err = errors.New("several errors were found. please check the log for more details")
  }

  return
}
