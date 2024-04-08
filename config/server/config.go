package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Transport    map[string]interface{} `yaml:"transport"`
	Storage      map[string]interface{} `yaml:"storage"`
	Registration bool                   `yaml:"registration"`
	Logger       map[string]interface{} `yaml:"log"`
}

func NewConfig(filepath string) (*Config, error) {
	cfg := &Config{}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	yamlDecoder := yaml.NewDecoder(f)
	if err = yamlDecoder.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
