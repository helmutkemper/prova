package commonData

import "strconv"

type DataFlightStretch struct {
	Origin      string
	Destination string
	Price       int
}

func (el *DataFlightStretch) PriceAsString() string {
	return strconv.Itoa(el.Price)
}
