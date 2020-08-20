package testDataSource

func (el *TestDataSource) deleteFlightStretch(origin, destination string) {
	el.mutex.Lock()
	defer el.mutex.Unlock()

	for k, flightStretch := range el.data {
		if flightStretch.Origin == origin && flightStretch.Destination == destination {
			el.data = append(el.data[:k-1], el.data[k:]...)
			return
		}
	}
}
