package flightSuggestion

import (
	"commonData"
	"dataSourceInterface"
	"errors"
)

func (el *FlightSuggestion) FindCheapestRoute(dataSource dataSourceInterface.DataSourceBasicInterface, origin, destination string) (cheapestRoute Route, err error) {
	if dataSource == nil {
		err = errors.New("data source not set")
		return
	}

	var tmpData = make([]commonData.DataFlightStretch, 0)
	var processingData FlightSuggestion
	var completeRoute FlightSuggestion
	var tmpRoutes FlightSuggestion

	tmpData, err = dataSource.GetFlightStretchByOriginIataCode(origin)
	if err != nil {
		return
	}

	for _, flightStretch := range tmpData {
		processingData.initNewRoute(flightStretch.Origin, flightStretch.Destination, flightStretch.Price)
	}

	for {
		tmpRoutes = processingData.popCompleteRoutes(origin, destination)
		completeRoute.addAnotherRouteList(&tmpRoutes)
		completeRoute.discardAllExpensiveRoutes()

		if processingData.getLength() == 0 {
			break
		}

		for tmpProcessingKey, tmpProcessingRoutes := range processingData.getList() {
			var lastDestination string
			var found bool
			lastDestination, found = tmpProcessingRoutes.getLastDestinationInThisRoute()
			if found == false {
				// neste ponto do código, uma rota não pode está sem dados
				err = errors.New("route.FindLowPriceRoute().error: bug")
				return
			}

			tmpData, _ = dataSource.GetFlightStretchByOriginIataCode(lastDestination)
			if len(tmpData) == 0 {
				processingData.deleteKey(tmpProcessingKey)
				continue
			}
			for _, newFlightStretchData := range tmpData {
				_ = processingData.copyRouteAndAddFlightStretch(&tmpProcessingRoutes, newFlightStretchData.Origin, newFlightStretchData.Destination, newFlightStretchData.Price)
			}
			processingData.deleteByReference(&tmpProcessingRoutes)
		}

	}

	var tmp = completeRoute.list
	if len(tmp) == 1 {
		cheapestRoute = tmp[0]
	}

	return
}
