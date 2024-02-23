package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	From     string   `yaml:"from"`
	To       []string `yaml:"to"`
	Password string   `yaml:"password"`
}

func ReadConfig(filePath string) Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return config
}
