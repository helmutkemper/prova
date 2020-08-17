package importFlightStretch

// Limpa a lista de erros.
// A lista contém todos os erros do arquivo processado
func (el *CSV) clearErrorList() {
	el.errorList = make([]error, 0)
}
