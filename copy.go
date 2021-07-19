package redisql

import (
	"context"

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

var ctx = context.Background()

// CopyToHash reads a desired SQL table's rows and writes them to Redis hashes
func (c Config) CopyToHash() error {
}

// CopyToString reads a desired SQL table's rows and writes them to Redis strings
func (c Config) CopyToString() error {
}

// CopyToList reads a desired SQL table's rows and writes them to Redis lists
func (c Config) CopyToList() error {
}
