package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTPPort          string `yaml:"httpPort"`
	GrpcServerAddress string `yaml:"grpcServerAddress"`
}

func NewConfig() (*Config, error) {
	var config Config

	pathToYamlFile := "config/config.yml"

	configFile, err := os.Open(pathToYamlFile)
	if err != nil {
		return nil, fmt.Errorf("error decode config file: %w", err)
	}

	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal yaml config: %w", err)
	}

	return &config, nil
}
