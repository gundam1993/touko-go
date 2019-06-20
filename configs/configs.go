package configs

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DBPath string `yaml:"db_path"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

var config *Config

func LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var configDate Config
	err = yaml.Unmarshal(data, &configDate)
	if err != nil {
		return err
	}
	config = &configDate
	return err
}

func GetConfig() *Config {
	return config
}
