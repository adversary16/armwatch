package main

import (
	"fmt"
	"strconv"
	"sync"

	"radiozi.ga/armwatch/cmd/battery"
	"radiozi.ga/armwatch/util/confParser"
	"radiozi.ga/armwatch/util/httpServer"
	"radiozi.ga/armwatch/util/repoize"
	"radiozi.ga/armwatch/util/routinify"
)

var configuration interface{}
var httpConfig httpServer.ServerSettings

func Init() {
	// initialize internal storage

	confParser.Init(&configuration)
	fmt.Println(configuration)
	httpConfig.Port = 8080
	httpConfig.Routes = httpServer.RouteMap{
		"/status": func(s string) (string, error) {
			return strconv.Itoa(battery.Get()), nil
		},
	}

	repoize.Init()
	// initialize data sources
	battery.Init()
	httpServer.Init(httpConfig)
}

func main() {
	var wg sync.WaitGroup
	Init()

	routinify.Add([]func() (int, error){
		battery.UpdateCapacity,
	}, &wg, 1000)

	wg.Wait()
	fmt.Println("over")
}
