package dataSourceInterface

import "commonData"

type DataSourceInterface interface{
  ClearAllData()
  AppendData(string, string, int)
  GetAllData() (*[]commonData.DataSegment, error)
  GetSegmentBySourceIataCode(string) ([]commonData.DataSegment, error)
  GetSegmentByDestinationIataCode(string) ([]commonData.DataSegment, error)
}
