package flightSuggestion

import (
	"commonData"
	"dataSourceInterface"
	"errors"
)

// Procura por rotas de voo mais baratas entre dois aeroportos
//   dataSource:    fonte de dados compatível com a interface dataSourceInterface.DataSourceBasicInterface contendo os trechos de voo e preços
//   origin:        aeroporto de origem no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   destination:   aeroporto de destino no formato IATA, um código de três letras maiúsculas usado para definir os aeroportos pelo mundo
//   cheapestRoute: lista de trechos da rota de voo mais barata com os preços individuais e o preço total
//   err:           erro no padrão Golang
//
// Esse foi um código simples, porém, bem trabalhoso de ser feito e como não faz parte
// dos códigos que sou acostumado a fazer, o meu modo de fazer é simples, fiz um código
// ruim e fácil de ser pensado para uma primeira tentativa e depois separei os módulos em
// funções. Para uma segunda entrega ele pode ser otimizado para consumo de memória, e
// algum outro problema, mas, minha filosofia é: esta é a primeira entrega e ela atende
// o que foi pedido. Se haverá uma melhora posterior, cabe ao product owner decidir.
//
// Caso a regra mude, a função completeRoute.discardAllExpensiveRoutes() é a única parte
// responsável pela regra, ou seja, funções pequenas de responsabilidade única.
//
// Como regra das empresas onde trabalhei, panic() é usado para gerar alarmes no servidor
// em caso de erros graves
func (el *FlightSuggestion) FindCheapestRoute(
	dataSource dataSourceInterface.DataSourceBasicInterface,
	origin,
	destination string,
) (
	cheapestRoute Route,
	err error,
) {

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
		tmpRoutes, err = processingData.popCompleteRoutes(origin, destination)
		if err != nil {
			err = errors.New("Route.FindLowPriceRoute().error: houston we have a problem! we have a bug")
			panic(err)
			return
		}
		completeRoute.addAnotherRouteList(&tmpRoutes)
		err = completeRoute.discardAllExpensiveRoutes()
		if err != nil {
			err = errors.New("Route.FindLowPriceRoute().error: houston we have a problem! we have a bug")
			panic(err)
			return
		}

		if processingData.getLength() == 0 {
			break
		}

		for tmpProcessingKey, tmpProcessingRoutes := range processingData.getList() {
			var lastDestination string
			var found bool
			lastDestination, found = tmpProcessingRoutes.getLastDestinationInThisRoute()
			if found == false {
				// neste ponto do código, uma rota não pode está sem dados
				err = errors.New("Route.FindLowPriceRoute().error: houston we have a problem! we have a bug")
				panic(err)
				return
			}

			tmpData, _ = dataSource.GetFlightStretchByOriginIataCode(lastDestination)
			if len(tmpData) == 0 {
				err = processingData.deleteKey(tmpProcessingKey)
				if err != nil {
					err = errors.New("Route.FindLowPriceRoute().error: houston we have a problem! we have a bug")
					panic(err)
				}
				continue
			}
			for _, newFlightStretchData := range tmpData {
				err = processingData.copyRouteAndAddFlightStretch(&tmpProcessingRoutes, newFlightStretchData.Origin, newFlightStretchData.Destination, newFlightStretchData.Price)
				if err != nil {
					err = errors.New("Route.FindLowPriceRoute().error: houston we have a problem! we have a bug")
					panic(err)
				}
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
