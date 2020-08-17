package testDataSource

import "commonData"

func (el *TestDataSource) GetAllData() (data *[]commonData.DataFlightStretch, err error) {
	el.mutex.Lock()
	defer el.mutex.Unlock()

	return &el.data, nil
}
