package importFlightStretch

import "commonData"

func (el CSV) dataToString(lineData commonData.DataSegment) string {
  return lineData.Origin+
    el.fieldSeparator+
    lineData.Destination+
    el.fieldSeparator+
    lineData.PriceAsString()
}
