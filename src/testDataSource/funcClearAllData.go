package testDataSource

import "commonData"

func (el *TestDataSource) ClearAllData() {
	el.mutex.Lock()
	defer el.mutex.Unlock()

	el.data = make([]commonData.DataFlightStretch, 0)
}
