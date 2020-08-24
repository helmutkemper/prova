package testDataSource

func (el *TestDataSource) GetEventChannel() (channel *chan bool) {
	return &el.Event
}
