package testDataSource

import (
	"commonData"
	"errors"
)

func (el *TestDataSource) GetSegmentBySourceIataCode(value string) (data []commonData.DataSegment, err error) {
	data = make([]commonData.DataSegment, 0)

	for _, line := range el.data {
		if line.Origin == value {
			data = append(data, line)
		}
	}

	if len(data) == 0 {
		err = errors.New("data not found")
	}

	return
}
