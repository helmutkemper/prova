package restServer

import (
	"encoding/json"
	"flightSuggestion"
	"net/http"
)

func (el apiRestful) setOutput(w http.ResponseWriter, route flightSuggestion.Route, err error) {
	var out output
	out.err = err
	out.route = route

	var data []byte
	data, _ = json.Marshal(&out)
	_, _ = w.Write(data)
}
