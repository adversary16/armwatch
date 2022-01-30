package dockerService

import (
	"context"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ContainerDTO struct {
	State string
	Name  string
	Id    string
}

func List() []ContainerDTO {
	var containerList []ContainerDTO
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Println(err)
	} else {
		containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		for _, container := range containers {
			containerDto := ContainerDTO{container.State, strings.Join(container.Names, ", "), container.ID}
			containerList = append(containerList, containerDto)
		}
	}
	return containerList
}
