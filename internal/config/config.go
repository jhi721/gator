package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilepath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homedir, configName), nil
}

func (c *Config) SetUser(userName string) error {
	configFilepath, err := getConfigFilepath()
	if err != nil {
		return err
	}

	c.CurrentUserName = userName

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFilepath, data, 0644); err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	configFilepath, err := getConfigFilepath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFilepath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
