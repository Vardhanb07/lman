package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	ErrInvalidConfig    = errors.New("invalid config")
	ErrConfigNotFound   = errors.New("config file not found")
	ErrPermissionDenied = errors.New("premission defined reading config file")
	ErrMissingFields    = errors.New("validation error: 'filepath' and 'linkpath' are required for all links")
)

type Link struct {
	Filepath string `mapstructure:"filepath"`
	Linkpath string `mapstructure:"linkpath"`
}

type Config struct {
	Links []Link `mapstructure:"links"`
}

func readConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrConfigNotFound
		}
		if errors.Is(err, os.ErrPermission) {
			return nil, ErrPermissionDenied
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidConfig, err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidConfig, err)
	}
	for _, links := range cfg.Links {
		if links.Filepath == "" || links.Linkpath == "" {
			return nil, ErrMissingFields
		}
	}
	return &cfg, nil
}
