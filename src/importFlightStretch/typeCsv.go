package importFlightStretch

import (
	"dataSourceInterface"
	"sync"
)

// Este pacote importa arquivos CSV com trechos de voos e valores no seguinte formato:
//   IATA code da origem;
//   IATA cede do destino;
//   Preço;
//
// Exemplo:
//   GRU,BRC,10
//   BRC,SCL,5
//   GRU,CDG,75
//
// Atenção:
//   Não foi solicitado que o preço fosse em ponto flutuante e nem levasse em conta a
//   conversão de moedas
type CSV struct {
	errorList      []error
	fieldSeparator string
	dataSource     dataSourceInterface.DataSourceBasicInterface
	mutex          sync.Mutex
}
