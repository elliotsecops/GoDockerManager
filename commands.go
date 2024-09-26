// commands.go
package main

import (
	"context"
		"fmt"
			"log"

				"github.com/docker/docker/api/types"
					"github.com/docker/docker/client"
						"github.com/spf13/cobra"
						)

						var listCmd = &cobra.Command{
							Use:   "list",
								Short: "List all running containers",
									Long:  `List detailed information about all running containers.`,
										Run: func(cmd *cobra.Command, args []string) {
												cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
														if err != nil {
																	log.Fatalf("Error creating Docker client: %v", err)
																			}
																					defer cli.Close()

																							containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
																									if err != nil {
																												log.Fatalf("Error listing containers: %v", err)
																														}

																																for _, container := range containers {
																																			fmt.Printf("ID: %s, Image: %s, Command: %s, Created: %s, Status: %s, Ports: %v\n",
																																							container.ID[:10], container.Image, container.Command, container.Created, container.Status, container.Ports)
																																									}
																																										},
																																										}

																																										var startCmd = &cobra.Command{
																																											Use:   "start [container ID]",
																																												Short: "Start a container",
																																													Long:  `Start a Docker container by its ID.`,
																																														Args:  cobra.ExactArgs(1),
																																															Run: func(cmd *cobra.Command, args []string) {
																																																	containerID := args[0]
																																																			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
																																																					if err != nil {
																																																								log.Fatalf("Error creating Docker client: %v", err)
																																																										}
																																																												defer cli.Close()

																																																														err = cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
																																																																if err != nil {
																																																																			log.Fatalf("Error starting container: %v", err)
																																																																					}

																																																																							fmt.Printf("Container %s started successfully.\n", containerID)
																																																																								},
																																																																								}

																																																																								var stopCmd = &cobra.Command{
																																																																									Use:   "stop [container ID]",
																																																																										Short: "Stop a container",
																																																																											Long:  `Stop a Docker container by its ID.`,
																																																																												Args:  cobra.ExactArgs(1),
																																																																													Run: func(cmd *cobra.Command, args []string) {
																																																																															containerID := args[0]
																																																																																	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
																																																																																			if err != nil {
																																																																																						log.Fatalf("Error creating Docker client: %v", err)
																																																																																								}
																																																																																										defer cli.Close()

																																																																																												err = cli.ContainerStop(context.Background(), containerID, nil)
																																																																																														if err != nil {
																																																																																																	log.Fatalf("Error stopping container: %v", err)
																																																																																																			}

																																																																																																					fmt.Printf("Container %s stopped successfully.\n", containerID)
																																																																																																						},
																																																																																																						}

																																																																																																						func init() {
																																																																																																							rootCmd.AddCommand(listCmd)
																																																																																																								rootCmd.AddCommand(startCmd)
																																																																																																									rootCmd.AddCommand(stopCmd)
																																																																																																									}