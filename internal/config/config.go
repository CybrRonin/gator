package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	// Updated to a more typical and scalable "read JSON from file" pattern
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	/*
	    * Original version, retained for reference
	   	data, err := os.ReadFile(path)
	   	if err != nil {
	   		return Config{}, err
	   	}

	   	cfg := Config{}
	   	err = json.Unmarshal(data, &cfg)
	   	if err != nil {
	   		return Config{}, err
	   	}
	*/
	return cfg, nil
}

func (cfg Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	return write(cfg)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("error retrieving home directory")
	}

	path := filepath.Join(home, configFileName)

	return path, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Updated to a more traditional/scalable style
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	/*
	    * Original version held for reference
	   	data, err := json.Marshal(cfg)
	   	if err != nil {
	   		return err
	   	}

	   	err = os.WriteFile(path, data, 0644)
	   	if err != nil {
	   		return err
	   	}
	*/
	return nil
}
