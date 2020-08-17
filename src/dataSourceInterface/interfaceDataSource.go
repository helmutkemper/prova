package dataSourceInterface

import "commonData"

type DataSourceBasicInterface interface {
	ClearAllData()
	AppendData(string, string, int)
	GetAllData() (*[]commonData.DataFlightStretch, error)
	GetSegmentBySourceIataCode(string) ([]commonData.DataFlightStretch, error)
	GetSegmentByDestinationIataCode(string) ([]commonData.DataFlightStretch, error)
}
