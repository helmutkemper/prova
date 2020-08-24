package dataSourceInterface

import "commonData"

type DataSourceBasicInterface interface {
	// apaga todos os dados da fonte de dados
	ClearAllData()

	// adiciona um novo trecho
	AppendData(source, destination string, price commonData.Price)

	// retorna toda a base de dados
	GetAllData() (data *[]commonData.DataFlightStretch, err error)

	// retorna a lista de trechos baseada na origem do voo
	GetFlightStretchByOriginIataCode(value string) (data []commonData.DataFlightStretch, err error)

	// pega todos os trechos baseados no destino
	GetFlightStretchByDestinationIataCode(value string) (data []commonData.DataFlightStretch, err error)

	// channel de evento quando um novo dado é adicionado
	GetEventChannel() (channel *chan bool)
}
