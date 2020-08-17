package testDataSource

import (
	"commonData"
	"errors"
)

func (el *TestDataSource) GetSegmentByDestinationIataCode(value string) (data []commonData.DataSegment, err error) {
	data = make([]commonData.DataSegment, 0)

	for _, line := range el.data {
		if line.Destination == value {
			data = append(data, line)
		}
	}

	if len(data) == 0 {
		err = errors.New("data not found")
	}

	return
}
