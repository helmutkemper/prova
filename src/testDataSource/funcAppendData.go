package testDataSource

import "commonData"

func (el *TestDataSource) AppendData(source, destination string, price int) {
	el.mutex.Lock()
	defer el.mutex.Unlock()

	if len(el.data) == 0 {
		el.data = make([]commonData.DataFlightStretch, 0)
	}

	el.data = append(el.data, commonData.DataFlightStretch{
		Origin:      source,
		Destination: destination,
		Price:       commonData.Price(price),
	})
}
