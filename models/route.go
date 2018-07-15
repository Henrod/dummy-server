package models

import "net/http"

// Route is a server route
type Route struct {
	Method   string
	Path     string
	Response *Response
}

// Handler returns the handler for route
func (r *Route) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(r.Response.Code)
		w.Write([]byte(r.Response.Body))
	}
}
