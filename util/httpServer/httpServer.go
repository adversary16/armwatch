package httpServer

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type RouteMap map[string]func(w http.ResponseWriter, r *http.Request)

type ServerSettings struct {
	Port int
	Host string
}

var predefinedRoutes = RouteMap{}

func ParseRoutes(routeMap RouteMap) error {
	for k, v := range routeMap {
		predefinedRoutes[k] = v
	}
	for path, handleFunc := range predefinedRoutes {
		http.HandleFunc(path, handleFunc)
	}
	return nil
}

func Init(config ServerSettings, routes RouteMap, wg *sync.WaitGroup) {
	ParseRoutes(routes)
	serveAdress := strings.Join([]string{
		config.Host,
		strconv.Itoa(config.Port),
	}, ":")
	go func() {
		log.Fatal(http.ListenAndServe(serveAdress, nil))
	}()
}
