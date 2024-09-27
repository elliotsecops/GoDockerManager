// internal/docker/commands.go
package docker

import (
	"fmt"
)

// ListContainersCommand lists all Docker containers using the provided DockerClient.
func ListContainersCommand(client DockerClient) {
	containers, err := client.ListContainers()
	if err != nil {
		fmt.Println("Error listing containers:", err)
		return
	}
	for _, container := range containers {
		fmt.Println(container)
	}
}

// StartContainerCommand starts a Docker container using the provided DockerClient.
func StartContainerCommand(client DockerClient, containerID string) {
	err := client.StartContainer(containerID)
	if err != nil {
		fmt.Println("Error starting container:", err)
		return
	}
	fmt.Println("Container started:", containerID)
}

// StopContainerCommand stops a Docker container using the provided DockerClient.
func StopContainerCommand(client DockerClient, containerID string) {
	err := client.StopContainer(containerID)
	if err != nil {
		fmt.Println("Error stopping container:", err)
		return
	}
	fmt.Println("Container stopped:", containerID)
}
