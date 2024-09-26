package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DockerSocket string `yaml:"docker_socket"` // Docker socket path
	OutputFormat string `yaml:"output_format"` // Output format (e.g., "table", "json")
}

var Cfg Config

func LoadConfig(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
