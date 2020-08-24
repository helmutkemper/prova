package restServer

import (
	"dataSourceInterface"
	"flightSuggestion"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

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
			api: api,
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
