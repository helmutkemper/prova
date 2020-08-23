package restServer

import (
	"commonData"
	"dataSourceInterface"
	"encoding/json"
	"flightSuggestion"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type output struct {
	Error   []string
	Success bool
	Length  int
	Object  []object
	route   flightSuggestion.Route
	err     error
}

type object struct {
	Price commonData.Price
	Route string
}

func (el output) MarshalJSON() (data []byte, err error) {
	if el.err != nil {

		el.Error = make([]string, 1)
		el.Error[0] = el.err.Error()

		el.Object = nil
		return json.Marshal(struct {
			Error   []string
			Success bool
			Length  int
			Object  []object
			route   flightSuggestion.Route
			err     error
		}{
			Error:  el.Error,
			Object: []object{},
		})
	}

	el.Success = true
	el.Length = 1
	el.Object = make([]object, 1)
	el.Object[0].Route = el.route.GetFlightStretchRouteAsString()
	el.Object[0].Price = el.route.GetPrice()

	return json.Marshal(struct {
		Error   []string
		Success bool
		Length  int
		Object  []object
		route   flightSuggestion.Route
		err     error
	}{
		Error:   []string{},
		Success: el.Success,
		Length:  el.Length,
		Object:  el.Object,
	})
}

type apiRestful struct {
	flightSuggestion *flightSuggestion.FlightSuggestion
	route            flightSuggestion.Route
	dataSource       dataSourceInterface.DataSourceBasicInterface
}

func (el apiRestful) setHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func (el apiRestful) setOutput(w http.ResponseWriter, route flightSuggestion.Route, err error) {
	var out output
	out.err = err
	out.route = route

	var data []byte
	data, _ = json.Marshal(&out)
	_, _ = w.Write(data)
}

type apiFlightSuggestion struct {
	Api apiRestful
}

func (el apiFlightSuggestion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	el.Api.setHeader(w)
	var err error
	var origin = r.URL.Query().Get("org")
	origin = strings.ToUpper(origin)

	var destination = r.URL.Query().Get("des")
	destination = strings.ToUpper(destination)

	err = el.Api.route.TestRules(origin, destination)
	if err != nil {
		el.Api.setOutput(w, el.Api.route, err)
		return
	}

	el.Api.route, err = el.Api.flightSuggestion.FindCheapestRoute(origin, destination)
	if err != nil {
		el.Api.setOutput(w, el.Api.route, err)
		return
	}
	el.Api.setOutput(w, el.Api.route, err)
}

type apiRoutAddToDataServer struct {
	Api         apiRestful
	Origin      string
	Destination string
	Price       commonData.Price
}

func (el apiRoutAddToDataServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	el.Api.setHeader(w)
	var err error
	var body []byte

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		el.Api.setOutput(w, el.Api.route, err)
		return
	}

	err = r.Body.Close()
	if err != nil {
		el.Api.setOutput(w, el.Api.route, err)
		return
	}

	err = json.Unmarshal(body, &el)
	if err != nil {
		el.Api.setOutput(w, el.Api.route, err)
		return
	}

	err = el.Api.route.Add(el.Origin, el.Destination, el.Price)
	if err != nil {
		el.Api.setOutput(w, el.Api.route, err)
		return
	}
	el.Api.dataSource.AppendData(el.Origin, el.Destination, el.Price)

	el.Api.setOutput(w, el.Api.route, err)
}

func Server(route flightSuggestion.Route, flightSuggestion *flightSuggestion.FlightSuggestion, dataSource dataSourceInterface.DataSourceBasicInterface, port int) {
	var api = apiRestful{
		route:            route,
		flightSuggestion: flightSuggestion,
		dataSource:       dataSource,
	}

	mux := http.NewServeMux()
	mux.Handle(
		"/api/find",
		apiFlightSuggestion{
			Api: api,
		},
	)
	mux.Handle(
		"/api/add",
		apiRoutAddToDataServer{
			Api: api,
		},
	)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		_, _ = fmt.Fprintf(w, "Welcome to the super economical airlines")
	})
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), mux))
}
