package flightSuggestion

import (
	"commonData"
	"errors"
	"math"
)

func (el *FlightSuggestion) calculatesTheCheapestRouteKey() (cheapestKey int, err error) {
	if len(el.list) == 0 {
		err = errors.New("route list is empty")
		return
	}

	var priceList = make([]commonData.Price, 0)
	for _, r := range el.list {
		priceList = append(priceList, r.GetPrice())
	}

	var cheapest = commonData.Price(math.MaxInt32)
	for k, price := range priceList {
		if cheapest > price {
			cheapestKey = k
			cheapest = price
		}
	}

	return
}
