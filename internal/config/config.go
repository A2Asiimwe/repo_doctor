package config

import (
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	GitHubToken string   `yaml:"github_token"`
	Repos       []string `yaml:"repos"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}