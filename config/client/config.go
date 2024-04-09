package client

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server string `yaml:"server"`
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
