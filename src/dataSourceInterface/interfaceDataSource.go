package dataSourceInterface

import "commonData"

type DataSourceInterface interface {
	ClearAllData()
	AppendData(string, string, int)
	GetAllData() (*[]commonData.DataFlightStretch, error)
	GetSegmentBySourceIataCode(string) ([]commonData.DataFlightStretch, error)
	GetSegmentByDestinationIataCode(string) ([]commonData.DataFlightStretch, error)
}
