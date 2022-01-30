package dockerService

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var dockerClient *client.Client

type ContainerDTO struct {
	State string
	Name  string
	Id    string
}

func GetClient() *client.Client {
	if dockerClient == nil {
		cli, err := client.NewClientWithOpts()
		if err != nil {
			panic(err)
		}
		dockerClient = cli
	}
	return dockerClient
}

func List() []ContainerDTO {
	var containerList []ContainerDTO
	cli := GetClient()
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	for _, container := range containers {
		containerDto := ContainerDTO{container.State, strings.Join(container.Names, ", "), container.ID}
		containerList = append(containerList, containerDto)
	}
	return containerList
}

func WatchEvents() {
	cli := GetClient()
	messageChan, _ := cli.Events(context.Background(), types.EventsOptions{})
	m := <-messageChan
	fmt.Println(m)
}
