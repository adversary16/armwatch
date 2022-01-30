package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var host http.Server

type RouteMap map[string]func(w http.ResponseWriter, r *http.Request)

type ServerSettings struct {
	Port int
	Host string
}

func basicReponse(data string) (string, error) {
	fmt.Println(data)
	return data, nil
}

var predefinedRoutes = RouteMap{}

func ParseRoutes(routeMap RouteMap) error {
	for k, v := range routeMap {
		predefinedRoutes[k] = v
	}
	for path, handleFunc := range predefinedRoutes {
		http.HandleFunc(path, handleFunc)
	}
	fmt.Println(len(predefinedRoutes), "routes initialized")
	return nil
}

func Init(config ServerSettings, routes RouteMap) *http.Server {
	host = http.Server{}
	ParseRoutes(routes)
	serveAdress := strings.Join([]string{
		config.Host,
		strconv.Itoa(config.Port),
	}, ":")
	log.Fatal(http.ListenAndServe(serveAdress, nil))
	return &host
}
