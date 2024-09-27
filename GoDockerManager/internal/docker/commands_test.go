// internal/docker/commands_test.go
package docker

import (
	"testing"
)

func TestListContainersCommand(t *testing.T) {
	mockClient := new(MockDockerClient)
	mockClient.On("ListContainers").Return([]string{"container1", "container2"}, nil)

	ListContainersCommand(mockClient)

	mockClient.AssertExpectations(t)
}

func TestStartContainerCommand(t *testing.T) {
	mockClient := new(MockDockerClient)
	mockClient.On("StartContainer", "container1").Return(nil)

	StartContainerCommand(mockClient, "container1")

	mockClient.AssertExpectations(t)
}

func TestStopContainerCommand(t *testing.T) {
	mockClient := new(MockDockerClient)
	mockClient.On("StopContainer", "container1").Return(nil)

	StopContainerCommand(mockClient, "container1")

	mockClient.AssertExpectations(t)
}
