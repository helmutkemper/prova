package testDataSource

import "commonData"

func (el *TestDataSource) GetAllData() (data *[]commonData.DataSegment, err error) {
	return &el.data, nil
}
