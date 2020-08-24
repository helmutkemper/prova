package commonData

import "strconv"

func (el *DataFlightStretch) PriceAsString() string {
	return strconv.Itoa(int(el.Price))
}
