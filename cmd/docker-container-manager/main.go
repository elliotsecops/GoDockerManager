package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/elliotsecops/GoDockerManager/internal/config"
	"github.com/elliotsecops/GoDockerManager/internal/docker"
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

var configFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "configs/config.yaml", "config file (default is configs/config.yaml)")

	// Load configuration from config.yaml (Optional - Remove if not used)
	config.LoadConfig(configFile)

	// Add commands here
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all running containers",
		Long:  `List detailed information about all running containers.`,
		Run: func(cmd *cobra.Command, args []string) {
			docker.ListContainersCommand()
		},
	}
	rootCmd.AddCommand(listCmd)

	startCmd := &cobra.Command{
		Use:   "start [container ID]",
		Short: "Start a container",
		Long:  `Start a Docker container by its ID.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			containerID := args[0]
			docker.StartContainerCommand(containerID)
		},
	}
	rootCmd.AddCommand(startCmd)

	stopCmd := &cobra.Command{
		Use:   "stop [container ID]",
		Short: "Stop a container",
		Long:  `Stop a Docker container by its ID.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			containerID := args[0]
			docker.StopContainerCommand(containerID)
		},
	}
	rootCmd.AddCommand(stopCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
