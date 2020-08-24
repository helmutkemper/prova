package restServer

import "flightSuggestion"

type output struct {
	Error   []string
	Success bool
	Length  int
	Object  []object
	route   flightSuggestion.Route
	err     error
}
