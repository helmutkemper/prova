package main

import (
	"errors"
	"factoryFlightStretch"
	"flightSuggestion"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

func main() {
	var err error
	var filePath string
	var wg sync.WaitGroup

	err = logToFileStart()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
	defer logToFileClose()

	filePath, err = findCsvFile()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	var flightSuggestionObj *flightSuggestion.FlightSuggestion
	flightSuggestionObj, err = factoryFlightStretch.NewWithTestDataSource(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("System successfully initialized\n")

	wg.Add(1)
	go func(flightSuggestionObj *flightSuggestion.FlightSuggestion) {
		var err error
		for {
			err = FindRouteByTerminal(flightSuggestionObj, &wg)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}(flightSuggestionObj)

	wg.Wait()
}

func FindRouteByTerminal(flightSuggestionObj *flightSuggestion.FlightSuggestion, wg *sync.WaitGroup) (err error) {
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

func findCsvFile() (filePath string, err error) {
	var fileName string

	if len(os.Args) == 1 {
		fmt.Println("Please enter the name of the route file:")
		_, err = fmt.Scan(&fileName)
		if err != nil {
			return
		}
	} else {
		fileName = os.Args[1]
		fmt.Printf("Loading file: %v\n", fileName)
	}

	filePath, err = fileFindRecursively(fileName)
	if err != nil {
		return
	}

	return
}

var logToFilePointer *os.File

func logToFileStart() (err error) {
	logToFilePointer, err = os.OpenFile("./logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logToFilePointer)
	return
}

func logToFileClose() {
	var err error

	err = logToFilePointer.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
}

func fileFindRecursively(fileName string) (filePath string, err error) {
	if _, err = os.Stat(fileName); os.IsNotExist(err) == false {
		filePath = fileName
		return
	}

	fileName = filepath.Base(fileName)
	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == fileName {
				filePath = path
				return nil
			}

			return nil
		},
	)

	if filePath == "" {
		err = errors.New("file not found")
	}

	return
}
