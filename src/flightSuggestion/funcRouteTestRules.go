package flightSuggestion

import "errors"

// testa todas as regras antes de adicionar um um trecho a uma rota
//   Atenção: não foi definida uma regra para preços de rotas ou gratitude da mesma
func (el *Route) TestRules(origin, destination string) (err error) {
	if el.testRulesGapsInRoute(origin) == true {
		err = errors.New("the addition of this flight stretch creates a gap")
		return
	}

	if el.testRulesRoutInfiniteLoop(destination) == true {
		err = errors.New("the addition of this flight stretch creates an infinite loop")
		return
	}

	err = el.testRulesCodeRules(origin, destination)
	return
}
