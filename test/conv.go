package test

import (
	"testing"

	redisql "github.com/DGKSK8LIFE/redisql"
)

func TestCopyToString(t *testing.T) {
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
