package main

import (
	"errors"

	"github.com/spf13/viper"
)

var ErrInvalidConfig = errors.New("Invalid config")

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
		return nil, ErrInvalidConfig
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, ErrInvalidConfig
	}
	return &cfg, nil
}
