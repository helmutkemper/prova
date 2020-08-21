package commonData

import "strconv"

type Price int

func (el Price) String() (price string) {
	return "$" + strconv.Itoa(int(el))
}

type DataFlightStretch struct {
	Origin      string
	Destination string
	Price       Price
}

func (el *DataFlightStretch) PriceAsString() string {
	return strconv.Itoa(int(el.Price))
}
