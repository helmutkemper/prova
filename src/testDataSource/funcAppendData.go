package testDataSource

import "commonData"

func (el *TestDataSource) AppendData(source, destination string, price int) {
	if len(el.data) == 0 {
		el.ClearAllData()
	}

	el.data = append(el.data, commonData.DataSegment{
		Origin:      source,
		Destination: destination,
		Price:       price,
	})
}
