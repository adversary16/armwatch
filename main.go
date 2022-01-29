package main

import (
	"fmt"

	"radiozi.ga/armwatch/cmd/battery"
)

func main() {
	battery.Init()
	fmt.Println(battery.Get())
}
