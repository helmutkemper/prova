package testDataSource

import (
  "commonData"
  "errors"
)

type TestDataSource struct {
  data []commonData.DataSegment
}

func(el *TestDataSource) ClearAllData() {
  el.data = make([]commonData.DataSegment, 0)
}

func(el *TestDataSource) AppendData(source, destination string, price int) {
  if len(el.data) == 0 {
    el.ClearAllData()
  }

  el.data = append(el.data, commonData.DataSegment{
    Origin:      source,
    Destination: destination,
    Price:       price,
  })
}

func(el *TestDataSource) GetAllData() (data *[]commonData.DataSegment, err error) {
  return &el.data, nil
}

func(el *TestDataSource) GetSegmentBySourceIataCode(value string) (data []commonData.DataSegment, err error) {
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

func(el *TestDataSource) GetSegmentByDestinationIataCode(value string) (data []commonData.DataSegment, err error){
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
