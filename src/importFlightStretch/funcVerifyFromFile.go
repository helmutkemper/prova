package importFlightStretch

import (
  "bufio"
  "errors"
  "os"
  "strconv"
)

func (el *CSV) verifyFromFile(filePointer *os.File) (err error) {
  var scanner *bufio.Scanner
  var lineCounter int
  var lineRawContent string
  var errorFoundList []error
  var origin string
  var destination string
  var price int

  scanner = bufio.NewScanner(filePointer)
  for scanner.Scan() {
    lineCounter += 1
    lineRawContent = scanner.Text()

    origin, destination, price, err = el.lineToData(lineRawContent)
    if err != nil {
      err = errors.New("csv processing file error on line "+strconv.Itoa(lineCounter)+" "+err.Error())
      el.errorList = append(el.errorList, err)
      continue
    }

    errorFoundList = el.verifyRules(origin, destination, price)
    if len(errorFoundList) != 0 {
      err = errors.New("csv processing file error on line "+strconv.Itoa(lineCounter))
      errorFoundList = append([]error{err}, errorFoundList...)
      el.errorList = append(el.errorList, errorFoundList...)
    }
  }

  if len(el.errorList) != 0 {
    err = errors.New("file has error")
  }
  return
}

