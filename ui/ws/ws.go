package ws

import "net/http"

var host http.Server
var routes map[string]func()

func ParseRoutes() error {
	// http.HandleFunc()
	return nil
}

func Init() *http.Server {
	host = http.Server{}
	host.ListenAndServe()
	return &host
}
