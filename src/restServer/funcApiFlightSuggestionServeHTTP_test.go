package restServer

import (
	"dataSourceInterface"
	"encoding/json"
	"factoryFlightStretch"
	"flightSuggestion"
	"io/ioutil"
	"net/http"
	"testing"
	"util"
)

func TestApiFlightSuggestion_ServeHTTP(t *testing.T) {
	var filePath = "input-file.correct.csv"

	var err error
	var serverBody []byte
	var serverData interface{}
	var response *http.Response
	var route *flightSuggestion.Route
	var flightSuggestionObj *flightSuggestion.FlightSuggestion
	var dataSource dataSourceInterface.DataSourceBasicInterface

	filePath, err = util.FileFindRecursively(filePath)
	if err != nil {
		t.Fail()
		panic(err)
	}

	route, flightSuggestionObj, dataSource, err = factoryFlightStretch.NewWithTestDataSourceAndAutoCsvSave(filePath)
	if err != nil {
		t.Fail()
		panic(err)
	}

	go func() {
		Server(*route, flightSuggestionObj, dataSource, 8080)
	}()

	response, err = http.Get("http://localhost:8080/api/find?org=BRC&des=CDG")
	if err != nil {
		t.Fatal(err)
	}
	serverBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(serverBody, &serverData)
	if err != nil {
		t.Fatal(err)
	}

	if serverData.(map[string]interface{})["Success"].(bool) != true {
		t.Fail()
	}

	if serverData.(map[string]interface{})["Length"].(float64) != 1 {
		t.Fail()
	}

	response, err = http.Get("http://localhost:8080/api/find?org=BRC&des=CD")
	if err != nil {
		t.Fatal(err)
	}
	serverBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(serverBody, &serverData)
	if err != nil {
		t.Fatal(err)
	}

	if serverData.(map[string]interface{})["Success"].(bool) != false {
		t.Fail()
	}

	if serverData.(map[string]interface{})["Length"].(float64) != 0 {
		t.Fail()
	}
}
