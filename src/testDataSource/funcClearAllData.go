package testDataSource

import "commonData"

func (el *TestDataSource) ClearAllData() {
	el.data = make([]commonData.DataFlightStretch, 0)
}
