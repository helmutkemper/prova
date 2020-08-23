package flightSuggestion

import (
	"errors"
	"regexp"
)

func (el *Route) testRulesCodeRules(origin, destination string) (err error) {
	var textRule *regexp.Regexp

	textRule, err = regexp.Compile("[A-Z]{3}")
	if err != nil {
		return
	}

	if len(origin) != 3 {
		err = errors.New("origin code must be three letters long")
		return
	}

	if textRule.MatchString(origin) == false {
		err = errors.New("origin code must be three letters long between 'A' and 'Z'")
		return
	}

	if len(destination) != 3 {
		err = errors.New("destination code must be three letters long")
		return
	}

	if textRule.MatchString(destination) == false {
		err = errors.New("destination code must be three letters long between 'A' and 'Z'")
		return
	}

	return
}
