package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config is a structure which represent configuretion for app
type Config struct {
	Port                    string `json:"port"`
	QuestionAnswersFilePath string `json:"QApath"`
	QuestionFilePath        string `json:"Qpath"`
	PlayersFilePath         string `json:"playersFilePath"`
}

// Load performs reading config file and creating config structure
func Load(path string) (config *Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
