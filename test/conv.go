package test

import (
	"testing"

	redisql "github.com/DGKSK8LIFE/redisql"
)

func TestCopyToString(t *testing.T) {
	/*
		- Needs to create database, and table schema from `schema.sql`
		- Then inserts random or preset data 1,000,000 times (1 million table rows)
		- Then runs CopyToString() on the table
		- Goals:
			1. Benchmark various Copy() functions
			2. Look for edgecases and major issues in them
	*/
	config := redisql.Config{
		SQLUser:     "root",
		SQLPassword: "password",
		SQLDatabase: "users",
		SQLTable:    "user",
		RedisAddr:   "localhost:6379",
		RedisPass:   "",
		Log:         true,
	}
	err := config.CopyToString()
	if err != nil {
		t.Error(err)
	}
}
