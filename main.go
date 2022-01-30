package main

import (
	"fmt"
	"sync"

	"radiozi.ga/armwatch/cmd/battery"
	"radiozi.ga/armwatch/ui/ws"
	"radiozi.ga/armwatch/util/confParser"
	"radiozi.ga/armwatch/util/httpServer"
	"radiozi.ga/armwatch/util/repoize"
	"radiozi.ga/armwatch/util/routinify"
)

var configuration confParser.Configuration
var httpConfig httpServer.ServerSettings

var wsRoutes = ws.WSRouteMap{
	"status": StatusController,
}

func Init(wg *sync.WaitGroup) {
	confParser.Init(&configuration)

	httpConf := configuration["http"].(map[string]interface{})
	httpConfig.Port = int(httpConf["port"].(int64))

	routes := httpServer.RouteMap{
		"/socket": ws.Controller(wsRoutes),
	}
	repoize.Init()
	// initialize data sources

	battery.Init()
	wg.Add(1)
	httpServer.Init(httpConfig, routes, wg)
}

func main() {
	wg := new(sync.WaitGroup)
	Init(wg)
	routinify.Add([]func() (int, error){
		battery.UpdateCapacity,
	}, 1000)

	fmt.Println("over")
	wg.Wait()
}
