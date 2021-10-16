package utils

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Config is the configuration struct for redisql
type Config struct {
	SQLType     string `yaml:"sqltype"`
	SQLUser     string `yaml:"sqluser"`
	SQLPassword string `yaml:"sqlpassword"`
	SQLDatabase string `yaml:"sqldatabase"`
	SQLHost     string `yaml:"sqlhost"`
	SQLPort     string `yaml:"sqlport"`
	SQLTable    string `yaml:"sqltable"`
	RedisAddr   string `yaml:"redisaddr"`
	RedisPass   string `yaml:"redispass"`
	LogLevel    *uint8 `yaml:"log_level"`
	LogFilename *string `yaml:"log_filename"`
}

// NewConfig initializes a new object of Config structure
func NewConfig(filePath string) (*Config, error) {
	if err := ValidateFilePath(filePath); err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var c Config
	if err = yaml.Unmarshal(file, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
