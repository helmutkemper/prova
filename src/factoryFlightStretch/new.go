package factoryFlightStretch

import (
  "dataSourceInterface"
  "importFlightStretch"
)

func New(dataSource dataSourceInterface.DataSourceInterface) (pointer *importFlightStretch.CSV) {
  pointer = &importFlightStretch.CSV{}
  pointer.SetDataSource(dataSource)
  pointer.SetFieldSeparator(",")

  return
}
