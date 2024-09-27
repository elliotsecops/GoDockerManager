package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ListContainers(cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func StartContainer(cli *client.Client, containerID string) error {
	return cli.ContainerStart(context.Background(), containerID, container.StartOptions{})
}

func StopContainer(cli *client.Client, containerID string) error {
	return cli.ContainerStop(context.Background(), containerID, container.StopOptions{})
}