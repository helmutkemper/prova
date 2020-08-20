package flightSuggestion

import "errors"

func (el *Route) testRules(origin, destination string) (err error) {
	if el.testRulesGapsInRoute(origin) == true {
		err = errors.New("the addition of this flight stretch creates a gap")
		return
	}

	if el.testRulesRoutInfiniteLoop(destination) == true {
		err = errors.New("the addition of this flight stretch creates an infinite loop")
		return
	}

	return
}
