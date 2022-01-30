package main

import (
	"log"

	"radiozi.ga/armwatch/cmd/battery"
)

type StatusRequest struct {
	payload string
}

type StatusResponse struct {
	BatteryCharge     int
	RunningContainers map[string]string
}

func StatusController(b []byte, jsonResponder func(interface{}) error) {
	response := StatusResponse{
		BatteryCharge: battery.Get(),
	}

	err := jsonResponder(response)
	if err != nil {
		log.Println(err)
	}
}
