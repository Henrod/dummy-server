package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Henrod/dummy-server/parser"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// App starts the server with routes
type App struct {
	router *mux.Router
}

// NewApp returns app by reading routes from parser
func NewApp(parser parser.Parser) (*App, error) {
	app := &App{}
	err := app.configure(parser)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) configure(parser parser.Parser) error {
	router := mux.NewRouter()

	routes, err := parser.Parse()
	if err != nil {
		return err
	}

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler()).Methods(route.Method)
	}

	a.router = router
	return nil
}

// Start starts listening on server
func (a *App) Start() {
	port := viper.GetInt("server.port")
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.router))
}
