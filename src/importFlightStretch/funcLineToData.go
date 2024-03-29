package importFlightStretch

import (
	"commonData"
	"errors"
	"strconv"
	"strings"
)

// converte uma linha CSV em dados para serem processados
func (el *CSV) lineToData(lineRawContent string) (origin, destination string, price commonData.Price, err error) {

	var priceInt64 int64
	var lineProcessedContent = strings.Split(lineRawContent, el.fieldSeparator)

	if len(lineProcessedContent) != 3 {
		err = errors.New("the file must contain three fields per line: IATA code from airport of origin with three letters; IATA code from destination airport with three letters; price of the flight stretch in numeric format")
		return
	}

	origin = lineProcessedContent[0]
	destination = lineProcessedContent[1]
	var priceString = lineProcessedContent[2]

	priceInt64, err = strconv.ParseInt(priceString, 10, 64)
	if err != nil {
		return
	}
	price = commonData.Price(priceInt64)

	return
}
