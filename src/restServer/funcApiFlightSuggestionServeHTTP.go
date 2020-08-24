package restServer

import (
	"net/http"
	"strings"
)

func (el apiFlightSuggestion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	el.api.setHeader(w)
	var err error
	var origin = r.URL.Query().Get("org")
	origin = strings.ToUpper(origin)

	var destination = r.URL.Query().Get("des")
	destination = strings.ToUpper(destination)

	err = el.api.route.TestRules(origin, destination)
	if err != nil {
		el.api.setOutput(w, el.api.route, err)
		return
	}

	el.api.route, err = el.api.flightSuggestion.FindCheapestRoute(origin, destination)
	if err != nil {
		el.api.setOutput(w, el.api.route, err)
		return
	}
	el.api.setOutput(w, el.api.route, err)
}
