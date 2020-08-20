package dataSourceInterface

import "commonData"

type DataSourceBasicInterface interface {
	ClearAllData()
	AppendData(string, string, commonData.Price)
	GetAllData() (*[]commonData.DataFlightStretch, error)
	GetFlightStretchByOriginIataCode(string) ([]commonData.DataFlightStretch, error)
	GetFlightStretchByDestinationIataCode(string) ([]commonData.DataFlightStretch, error)
}
