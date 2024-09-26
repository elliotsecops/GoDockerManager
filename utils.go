// utils.go
package main

import (
	"context"
		"fmt"
			"log"

				"github.com/docker/docker/api/types"
					"github.com/docker/docker/client"
					)

					// CreateDockerClient creates and returns a Docker client.
					func CreateDockerClient() (*client.Client, error) {
						cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
							if err != nil {
									return nil, fmt.Errorf("error creating Docker client: %v", err)
										}
											return cli, nil
											}

											// ListContainers lists all running containers.
											func ListContainers(cli *client.Client) ([]types.Container, error) {
												containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
													if err != nil {
															return nil, fmt.Errorf("error listing containers: %v", err)
																}
																	return containers, nil
																	}

																	// StartContainer starts a container by its ID.
																	func StartContainer(cli *client.Client, containerID string) error {
																		err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
																			if err != nil {
																					return fmt.Errorf("error starting container: %v", err)
																						}
																							return nil
																							}

																							// StopContainer stops a container by its ID.
																							func StopContainer(cli *client.Client, containerID string) error {
																								err := cli.ContainerStop(context.Background(), containerID, nil)
																									if err != nil {
																											return fmt.Errorf("error stopping container: %v", err)
																												}
																													return nil
																													}