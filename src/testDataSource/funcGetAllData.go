package testDataSource

import "commonData"

func (el *TestDataSource) GetAllData() (data *[]commonData.DataFlightStretch, err error) {
	return &el.data, nil
}
