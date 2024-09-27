// main.go
package main

import (
	"fmt"
	"os"

	"github.com/elliotsecops/GoDockerManager/internal/docker"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-container-manager",
	Short: "A simple Docker container management CLI",
	Long: `docker-container-manager is a CLI tool for managing Docker containers.
It allows you to list, start, and stop containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Docker Container Manager!")
		fmt.Println("Use --help for more information about available commands.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		client := docker.RealDockerClient{}
		docker.ListContainersCommand(&client)
	},
}

var startCmd = &cobra.Command{
	Use:   "start [containerID]",
	Short: "Start a Docker container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := docker.RealDockerClient{}
		docker.StartContainerCommand(&client, args[0])
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop [containerID]",
	Short: "Stop a Docker container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := docker.RealDockerClient{}
		docker.StopContainerCommand(&client, args[0])
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}