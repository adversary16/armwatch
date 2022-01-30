package main

import (
	"log"

	"radiozi.ga/armwatch/cmd/battery"
	"radiozi.ga/armwatch/cmd/dockerService"
)

type StatusRequest struct {
	payload string
}

type StatusResponse struct {
	BatteryCharge     int
	RunningContainers []dockerService.ContainerDTO
}

func StatusController(b []byte, jsonResponder func(interface{}) error) {
	response := StatusResponse{
		BatteryCharge:     battery.Get(),
		RunningContainers: dockerService.List(),
	}

	err := jsonResponder(response)
	if err != nil {
		log.Println(err)
	}
}
