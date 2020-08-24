package terminal

import (
	"errors"
	"flightSuggestion"
	"fmt"
	"os"
	"regexp"
)

func FindRouteByTerminal(flightSuggestionObj *flightSuggestion.FlightSuggestion) (err error) {
	var travelStringToProcess string
	var travelRegexp *regexp.Regexp
	var travelFlightStretch []string
	var travelFlightStretchStart string
	var travelFlightStretchEnd string

	var cheapestRoute flightSuggestion.Route

	fmt.Printf("\n\n")
	fmt.Printf("Where do you want to travel?\n")
	fmt.Printf("Please inform the airport of origin and destination\n")
	fmt.Printf("Example: GRU-CDG\n")
	_, err = fmt.Scan(&travelStringToProcess)
	if err != nil {
		return
	}

	if travelStringToProcess == "exit" {
		os.Exit(0)
	}

	travelRegexp = regexp.MustCompile(`\W`)
	travelFlightStretch = travelRegexp.Split(travelStringToProcess, -1)
	if len(travelFlightStretch) != 2 {
		err = errors.New("please, use origin code + '-' + destination code. example: GRU-CDG")
		return
	}

	travelFlightStretchStart = travelFlightStretch[0]
	travelFlightStretchEnd = travelFlightStretch[1]

	if len(travelFlightStretchStart) != 3 {
		err = errors.New("origin code must have three letters. example: GRU-CDG")
		return
	}
	if len(travelFlightStretchEnd) != 3 {
		err = errors.New("destination code must have three letters. example: GRU-CDG")
		return
	}
	if travelFlightStretchStart == travelFlightStretchEnd {
		err = errors.New("both codes are identical")
		return
	}

	cheapestRoute, err = flightSuggestionObj.FindCheapestRoute(travelFlightStretchStart, travelFlightStretchEnd)
	if err != nil {
		return
	} else {
		fmt.Printf("Cheapest route: %v\n\n", cheapestRoute)
	}

	return
}
