package models

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Route is a server route
type Route struct {
	Method   string
	Path     string
	Response *Response
}

// Handler returns the handler for route
func (r *Route) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Print(err.Error())
			return
		}

		log.Println("received request", string(body))
		w.WriteHeader(r.Response.Code)
		w.Write([]byte(r.Response.Body))
	}
}
