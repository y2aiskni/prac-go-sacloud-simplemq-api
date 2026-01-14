package app

import (
	"log"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Sacloud ConfigSacloud `yaml:"Sacloud"`
}

type ConfigSacloud struct {
	AccessToken       string         `yaml:"AccessToken"`
	AccessTokenSecret string         `yaml:"AccessTokenSecret"`
	Simplemq          ConfigSimplemq `yaml:"Simplemq"`
}

type ConfigSimplemq struct {
	ResourceID string `yaml:"ResourceID"`
	QueueName  string `yaml:"QueueName"`
	APIKey     string `yaml:"APIKey"`
}

func ReadConfig(path string) Config {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)

	}
	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		log.Fatalf("Failed to unmarshal: %s", err)
	}

	return config
}
