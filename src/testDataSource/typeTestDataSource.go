package testDataSource

import (
	"commonData"
	"sync"
)

type TestDataSource struct {
	data  []commonData.DataFlightStretch
	mutex sync.Mutex
}
