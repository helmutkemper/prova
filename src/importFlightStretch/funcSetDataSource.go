package importFlightStretch

import "dataSourceInterface"

func (el *CSV) SetDataSource(dataSource dataSourceInterface.DataSourceInterface) {
  el.dataSource = dataSource
}
