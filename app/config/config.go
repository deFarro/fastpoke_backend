package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Version string `yaml:"version"`
	AppPort string `yaml:"appPort"`
	DatabaseAddr string `yaml:"databaseAddr"`
	DatabasePort string `yaml:"databasePort"`
	DatabaseName string `yaml:"databaseName"`
	DatabaseUser string `yaml:"databaseUser"`
	DatabasePassword string `yaml:"databasePassword"`
}

func GetConfig(path string) (Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
