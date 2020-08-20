package testDataSource

import "testing"

// Quando escrevo um teste, e não há tempo de fazer cobertura total, eu sempre dou
// prioridade a escrever primeiro tudo aquilo que eu tenho alguma dúvida se está correto
// depois, tudo o que é crítico para o sistema, o que gera mais valor para o cliente e no
// final, aquilo sem muito valor para o cliente (algo pouco usado)
func TestDeleteFromSlice(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4}

	// apaga a primeira chave
	deleteKey := 0
	slice = append(slice[:deleteKey], slice[deleteKey+1:]...)
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 || slice[3] != 4 || len(slice) != 4 {
		t.Fail()
	}

	// apaga a última chave
	deleteKey = 3
	slice = append(slice[:deleteKey], slice[deleteKey+1:]...)
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 || len(slice) != 3 {
		t.Fail()
	}

	// apaga uma chave do meio
	deleteKey = 1
	slice = append(slice[:deleteKey], slice[deleteKey+1:]...)
	if slice[0] != 1 || slice[1] != 3 || len(slice) != 2 {
		t.Fail()
	}

	// apaga o resto por desencargo de consciência
	deleteKey = 0
	slice = append(slice[:deleteKey], slice[deleteKey+1:]...)
	if slice[0] != 3 || len(slice) != 1 {
		t.Fail()
	}

	deleteKey = 0
	slice = append(slice[:deleteKey], slice[deleteKey+1:]...)
	if len(slice) != 0 {
		t.Fail()
	}
}
