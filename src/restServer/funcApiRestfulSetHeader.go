package restServer

import "net/http"

func (el apiRestful) setHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
