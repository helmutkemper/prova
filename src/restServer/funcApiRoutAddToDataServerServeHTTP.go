package restServer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
