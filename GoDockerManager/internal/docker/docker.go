// internal/docker/docker.go
package docker

type DockerClient interface {
	ListContainers() ([]string, error)
	StartContainer(containerID string) error
	StopContainer(containerID string) error
}

type RealDockerClient struct{}

func (r *RealDockerClient) ListContainers() ([]string, error) {
	// Implementation to list containers
	return []string{"container1", "container2"}, nil
}

func (r *RealDockerClient) StartContainer(containerID string) error {
	// Implementation to start a container
	return nil
}

func (r *RealDockerClient) StopContainer(containerID string) error {
	// Implementation to stop a container
	return nil
}
