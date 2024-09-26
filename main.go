// main.go
package main

import (
	"fmt"
		"log"
			"os"

				"github.com/spf13/cobra"
					"gopkg.in/yaml.v3"
					)

					var rootCmd = &cobra.Command{
						Use:   "docker-container-manager",
							Short: "A CLI tool to manage Docker containers",
								Long: `A command-line tool written in Go to manage Docker containers. 
								This tool allows you to list, start, stop, restart, view logs, and execute commands inside containers.`,
									Run: func(cmd *cobra.Command, args []string) {
											// The root command doesn't do anything itself.
													// It just provides the base for other commands.
														},
														}

														var config Config

														func init() {
															// Load configuration from config.yaml
																loadConfig()

																	// Add commands here
																		rootCmd.AddCommand(listCmd)
																			rootCmd.AddCommand(startCmd)
																				rootCmd.AddCommand(stopCmd)
																				}

																				func loadConfig() {
																					// Read the config.yaml file
																						data, err := os.ReadFile("config.yaml")
																							if err != nil {
																									log.Fatalf("Error reading config file: %v", err)
																										}

																											// Unmarshal the YAML data into the Config struct
																												err = yaml.Unmarshal(data, &config)
																													if err != nil {
																															log.Fatalf("Error parsing config file: %v", err)
																																}
																																}

																																func main() {
																																	if err := rootCmd.Execute(); err != nil {
																																			fmt.Println(err)
																																					os.Exit(1)
																																						}
																																						}