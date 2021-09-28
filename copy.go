package redisql

import (
	"io/ioutil"

	"github.com/DGKSK8LIFE/redisql/utils"
	_ "github.com/go-sql-driver/mysql"
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
}

// NewConfig initializes a new object of Config structure
func NewConfig(filePath string) (*Config, error) {
	if err := utils.ValidateFilePath(filePath); err != nil {
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

// CopyToString reads a desired SQL table's rows and writes them to Redis strings
func (c Config) CopyToString() error {
	if err := utils.Convert("string", c.SQLUser, c.SQLPassword, c.SQLDatabase, c.SQLHost, c.SQLPort, c.SQLTable, c.RedisAddr, c.RedisPass, c.SQLType); err != nil {
		return err
	}
	return nil
}

// CopyToList reads a desired SQL table's rows and writes them to Redis lists
func (c Config) CopyToList() error {
	if err := utils.Convert("list", c.SQLUser, c.SQLPassword, c.SQLDatabase, c.SQLHost, c.SQLPort, c.SQLTable, c.RedisAddr, c.RedisPass, c.SQLType); err != nil {
		return err
	}
	return nil
}

// CopyToHash reads a desired SQL table's rows and writes them to Redis hashes
func (c Config) CopyToHash() error {
	if err := utils.Convert("hash", c.SQLUser, c.SQLPassword, c.SQLDatabase, c.SQLHost, c.SQLPort, c.SQLTable, c.RedisAddr, c.RedisPass, c.SQLType); err != nil {
		return err
	}
	return nil
}
