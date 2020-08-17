package importFlightStretch


import "dataSourceInterface"

type CSV struct {
  errorList      []error
  fieldSeparator string
  dataSource     dataSourceInterface.DataSourceInterface
}
