// Package config is responsible for reading and writing to the json file.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config filepath: %w", err)
	}
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %w", err)
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error during json unmarshal: %w", err)
	}
	return config, nil
}

func (c Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(c)
	if err != nil {
		return fmt.Errorf("error writing to disk: %w", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeLocation, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error finding home location: %w", err)
	}
	return filepath.Join(homeLocation, configFileName), nil
}

func write(config Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config filepath: %w", err)
	}
	file, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling json: %w", err)
	}
	err = os.WriteFile(filePath, file, 0o644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}
