package main

import (
	"flag"
	"os"

	redisql "github.com/DGKSK8LIFE/redisql"
)

var config redisql.Config

func init() {
	copyFlag := flag.NewFlagSet("copy", flag.ExitOnError)
	configFile := copyFlag.String("config", "", "yaml config file")
	copyFlag.Parse(os.Args[2:])
	config = redisql.Config{
		SQLUser:     *user,
		SQLPassword: *password,
		SQLDatabase: *database,
		SQLTable:    *table,
		RedisAddr:   *redisAddr,
		RedisPass:   *redisPass,
		Log:         *log,
	}
}

func main() {
	err := config.Copy()
	if err != nil {
		panic(err)
	}
}
