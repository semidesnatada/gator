package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigPath() (string, error) {

	homeDirPath, homePathErr := os.UserHomeDir()
	if homePathErr != nil {
		return "", homePathErr
	}

	const configFileName = "/.gatorconfig.json"
	path := homeDirPath + configFileName

	return path, nil
}

func Read() (Config, error) {

	path, pathErr := getConfigPath()
	if pathErr != nil {
		return Config{}, pathErr
	}

	data, readErr := os.ReadFile(path)
	if readErr != nil {
		return Config{}, readErr
	}

	var output Config
	if unmarshallErr := json.Unmarshal(data, &output); unmarshallErr != nil {
		return Config{}, unmarshallErr
	}
	return output, nil
}

func (c *Config) SetUser(userName string) error {
	if len(userName) == 0 {
		return errors.New("can't set a nil username")
	}
	c.CurrentUserName = userName

	writeErr := write(c)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func write(c *Config) error {
	toWrite, marshallErr := json.Marshal(c)
	if marshallErr != nil {
		return marshallErr
		// return errors.New("error marshalling config to json")
	}

	writePath, pathErr := getConfigPath()
	if pathErr != nil {
		return pathErr
	}

	writeErr := os.WriteFile(writePath, toWrite, 0664)
	if writeErr != nil {
		return writeErr
	}

	return nil
}