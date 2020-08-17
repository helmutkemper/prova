package importFlightStretch

import (
  "bufio"
  "os"
)

func (el *CSV) fileToDataSourceAppend(filePointer *os.File) (err error) {
  var scanner *bufio.Scanner
  var lineRawContent string
  var origin string
  var destination string
  var price int

  scanner = bufio.NewScanner(filePointer)
  for scanner.Scan() {
    lineRawContent = scanner.Text()
    origin, destination, price, err = el.lineToData(lineRawContent)
    if err != nil {
      return
    }

    el.dataSource.AppendData(origin, destination, price)
  }

  return
}

