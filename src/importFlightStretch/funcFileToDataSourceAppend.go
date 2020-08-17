package importFlightStretch

import (
	"bufio"
	"os"
)

// Adiciona os dados de uma arquivo ao data source após verificação de erro.
//   Atenção: Esta função não tem a responsabilidade de verificar dados repetidos.
func (el *CSV) fileToDataSourceAppend(filePointer *os.File) (err error) {
	var scanner *bufio.Scanner
	var lineRawContent string
	var origin string
	var destination string
	var price int

	scanner = bufio.NewScanner(filePointer)
	for scanner.Scan() {
		lineRawContent = scanner.Text()
		origin, destination, price, err = el.lineToData(lineRawContent)
		if err != nil {
			return
		}

		el.dataSource.AppendData(origin, destination, price)
	}

	return
}
