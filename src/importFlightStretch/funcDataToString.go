package importFlightStretch

import "commonData"

// Converte commonData.DataSegment em string estilo linha de CSV
func (el CSV) dataToString(lineData commonData.DataFlightStretch) string {
	return lineData.Origin +
		el.fieldSeparator +
		lineData.Destination +
		el.fieldSeparator +
		lineData.PriceAsString()
}
