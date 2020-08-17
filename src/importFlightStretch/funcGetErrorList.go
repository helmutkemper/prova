package importFlightStretch

// Retorna a lista de erros
func (el *CSV) GetErrorList() (errorList []error) {
	return el.errorList
}
