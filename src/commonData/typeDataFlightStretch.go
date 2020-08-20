package commonData

import "strconv"

type Price int

type DataFlightStretch struct {
	Origin      string
	Destination string
	Price       Price
}

func (el *DataFlightStretch) PriceAsString() string {
	return strconv.Itoa(int(el.Price))
}
