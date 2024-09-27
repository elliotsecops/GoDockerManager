// internal/docker/mock_docker.go
package docker

import (
	"github.com/stretchr/testify/mock"
)

type MockDockerClient struct {
	mock.Mock
}

func (m *MockDockerClient) ListContainers() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockDockerClient) StartContainer(containerID string) error {
	args := m.Called(containerID)
	return args.Error(0)
}

func (m *MockDockerClient) StopContainer(containerID string) error {
	args := m.Called(containerID)
	return args.Error(0)
}
