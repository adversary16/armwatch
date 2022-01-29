package main

import (
	"fmt"
	"sync"

	"radiozi.ga/armwatch/cmd/battery"
	"radiozi.ga/armwatch/util/routinify"
)

func Init() {
	battery.Init()
}

func main() {
	var wg sync.WaitGroup
	Init()
	routinify.Add([]func() (int, error){
		battery.UpdateCapacity}, &wg, 1000)
	wg.Wait()
	fmt.Println("over")
}
