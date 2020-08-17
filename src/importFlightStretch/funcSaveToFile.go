package importFlightStretch

import (
  "commonData"
  "errors"
  "os"
)

func (el *CSV) saveToFile(filePath string) (err error) {
  var filePointer *os.File
  var dataSource *[]commonData.DataSegment

  filePointer, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
  if err != nil {
    err = errors.New("csv save file error: "+err.Error())
    el.errorList = append(el.errorList, err)
    return
  }

  dataSource, err = el.dataSource.GetAllData()
  if err != nil {
    err = errors.New("csv save file error: "+err.Error())
    el.errorList = append(el.errorList, err)
    return
  }

  for _, lineData := range *dataSource {
    var line = el.dataToString(lineData)
    _, err = filePointer.WriteString(line)
    if err != nil {
      err = errors.New("csv save file error: "+err.Error())
      el.errorList = append(el.errorList, err)
      return
    }
  }

  return
}
