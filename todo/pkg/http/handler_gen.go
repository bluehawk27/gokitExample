// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http1 "net/http"

	endpoint "github.com/bluehawk27/gokitExample/todo/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

//  NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeGetHandler(m, endpoints, options["Get"])
	makeAddHandler(m, endpoints, options["Add"])
	makeSetCompleteHandler(m, endpoints, options["SetComplete"])
	makeRemoveCompleteHandler(m, endpoints, options["RemoveComplete"])
	makeDeleteHandler(m, endpoints, options["Delete"])
	return m
}
