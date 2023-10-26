package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Drivername     string `json:"drivername"`
		DataSourceName string `json:"dataSourceName"`
	}
	Port string `json:"port"`
}

func StartConfiguration() (*Config, error) {

	var config Config

	configFile, err := os.ReadFile("config/config.json")

	if err != nil {
		fmt.Println("Error reading the file", err)
		return nil, err
	}

	if err = json.Unmarshal(configFile, &config); err != nil {
		fmt.Println("Error converting to JSON", err)
		return nil, err
	}

	return &config, nil
}
