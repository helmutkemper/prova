package main

import (
	"dataSourceInterface"
	"errors"
	"factoryFlightStretch"
	"flightSuggestion"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"restServer"
	"sync"
	"terminal"
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

	//restServer.Server()

	filePath, err = findCsvFile()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	var route *flightSuggestion.Route
	var flightSuggestionObj *flightSuggestion.FlightSuggestion
	var dataSource dataSourceInterface.DataSourceBasicInterface
	route, flightSuggestionObj, dataSource, err = factoryFlightStretch.NewWithTestDataSource(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("System successfully initialized\n")

	wg.Add(1)

	terminal.New(flightSuggestionObj)
	go func() {
		restServer.Server(*route, flightSuggestionObj, dataSource, 8080)
	}()

	wg.Wait()
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
