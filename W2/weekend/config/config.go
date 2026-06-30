package config

import (
	"fmt"
	"os"
	"sync/atomic"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port         int    `yaml:"port"`
		ReadTimeout  string `yaml:"read_timeout"`
		WriteTimeout string `yaml:"write_timeout"`
	} `yaml:"server"`
	Services map[string]ServiceConf `yaml:"services"`
	LogLevel string                 `yaml:"log_level"`
}

type ServiceConf struct {
	URL      string `yaml:"url"`
	Critical bool   `yaml:"critical"`
}

var current atomic.Pointer[Config]

func Load() (*Config, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "/etc/w2/config.yaml"
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}
	var c Config
	if err := yaml.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	current.Store(&c)
	return &c, nil
}

func Get() *Config { return current.Load() }
