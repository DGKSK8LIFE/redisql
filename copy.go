package redisql

import (
	"github.com/DGKSK8LIFE/redisql/utils"
	_ "github.com/go-sql-driver/mysql"
)

// Configuration struct for redisql
type Config struct {
	SQLUser     string `yaml:"sqluser"`
	SQLPassword string `yaml:"sqlpassword"`
	SQLDatabase string `yaml:"sqldatabase"`
	SQLTable    string `yaml:"sqltable"`
	RedisAddr   string `yaml:"redisaddr"`
	RedisPass   string `yaml:"redispass"`
	Log         bool   `yaml:"log"`
}

// CopyToString reads a desired SQL table's rows and writes them to Redis strings
func (c Config) CopyToString() error {
	if err := utils.Convert("string", c.SQLUser, c.SQLPassword, c.SQLDatabase, c.SQLTable, c.RedisAddr, c.RedisPass, c.Log); err != nil {
		return err
	}
	return nil
}

// CopyToList reads a desired SQL table's rows and writes them to Redis lists
func (c Config) CopyToList() error {
	if err := utils.Convert("list", c.SQLUser, c.SQLPassword, c.SQLDatabase, c.SQLTable, c.RedisAddr, c.RedisPass, c.Log); err != nil {
		return err
	}
	return nil
}

// CopyToHash reads a desired SQL table's rows and writes them to Redis hashes
func (c Config) CopyToHash() error {
	if err := utils.Convert("hash", c.SQLUser, c.SQLPassword, c.SQLDatabase, c.SQLTable, c.RedisAddr, c.RedisPass, c.Log); err != nil {
		return err
	}
	return nil
}