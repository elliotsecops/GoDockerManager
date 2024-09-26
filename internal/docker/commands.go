package docker

import (
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

func ListContainersCommand() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Error creating Docker client:", err)
		os.Exit(1)
	}

	containers, err := ListContainers(cli)
	if err != nil {
		fmt.Println("Error listing containers:", err)
		os.Exit(1)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

func StartContainerCommand(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Error creating Docker client:", err)
		os.Exit(1)
	}

	err = StartContainer(cli, containerID)
	if err != nil {
		fmt.Println("Error starting container:", err)
		os.Exit(1)
	}

	fmt.Println("Container started successfully")
}

func StopContainerCommand(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Error creating Docker client:", err)
		os.Exit(1)
	}

	err = StopContainer(cli, containerID)
	if err != nil {
		fmt.Println("Error stopping container:", err)
		os.Exit(1)
	}

	fmt.Println("Container stopped successfully")
}