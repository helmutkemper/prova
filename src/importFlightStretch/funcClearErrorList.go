package importFlightStretch

func (el *CSV) clearErrorList() {
  el.errorList = make([]error, 0)
}
