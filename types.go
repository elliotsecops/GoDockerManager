// types.go
package main

// Config represents the configuration settings for the application.
type Config struct {
	DockerSocket string `yaml:"docker_socket"` // Docker socket path
		OutputFormat string `yaml:"output_format"` // Output format (e.g., "table", "json")
		}