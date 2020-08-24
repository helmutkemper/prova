package toolTerminal

import (
	"flightSuggestion"
	"fmt"
	"terminal"
)

func New(flightSuggestionObj *flightSuggestion.FlightSuggestion) {
	go func(flightSuggestionObj *flightSuggestion.FlightSuggestion) {
		var err error
		for {
			err = terminal.FindRouteByTerminal(flightSuggestionObj)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}(flightSuggestionObj)
}
