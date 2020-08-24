package restServer

import (
	"bytes"
	"dataSourceInterface"
	"encoding/json"
	"factoryFlightStretch"
	"flightSuggestion"
	"io/ioutil"
	"net/http"
	"testing"
	"util"
)

func TestApiRoutAddToDataServer_ServeHTTP(t *testing.T) {
	var filePath = "input-file.correct.csv"

	var err error
	var serverBody []byte
	var serverData interface{}
	var response *http.Response
	var route *flightSuggestion.Route
	var flightSuggestionObj *flightSuggestion.FlightSuggestion
	var dataSource dataSourceInterface.DataSourceBasicInterface
	var postBody []byte

	filePath, err = util.FileFindRecursively(filePath)
	if err != nil {
		t.Fail()
		panic(err)
	}

	route, flightSuggestionObj, dataSource, err = factoryFlightStretch.NewWithTestDataSource(filePath)
	if err != nil {
		t.Fail()
		panic(err)
	}

	go func() {
		Server(*route, flightSuggestionObj, dataSource, 8080)
	}()

	postBody = []byte(`{"origin":"REC","destination":"FLN","price":1}`)
	response, err = http.Post("http://localhost:8080/api/add", `Content-Type: application/json`, bytes.NewBuffer(postBody))
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
}
