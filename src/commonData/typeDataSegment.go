package commonData

import "strconv"

type DataSegment struct {
  Origin      string
  Destination string
  Price int
}

func (el *DataSegment) PriceAsString() string {
  return strconv.Itoa(el.Price)
}