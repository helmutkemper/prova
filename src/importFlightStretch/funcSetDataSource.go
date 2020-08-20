package importFlightStretch

import "dataSourceInterface"

// Recebe o ponteiro da fonte de dados
//
//   A fonte de dados deve conter os seguintes métodos públicos:
//
//   ClearAllData()
//   AppendData(string, string, int)
//   GetAllData() (*[]commonData.DataSegment, error)
//   GetFlightStretchBySourceIataCode(string) ([]commonData.DataSegment, error)
func (el *CSV) SetDataSource(dataSource dataSourceInterface.DataSourceBasicInterface) {
	el.dataSource = dataSource
}
