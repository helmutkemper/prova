package importFlightStretch

import "os"

func (el *CSV) ClearDataSourceAndLoadFile(filePath string) (err error) {
  var filePointer *os.File

  el.clearErrorList()
  filePointer, err = el.loadFile(filePath)
  if err != nil {
    return
  }

  err = el.verifyFromFile(filePointer)
  if err != nil {
    return
  }

  el.dataSource.ClearAllData()

  filePointer, err = el.loadFile(filePath)
  if err != nil {
    return
  }

  err = el.fileToDataSourceAppend(filePointer)

  return
}

