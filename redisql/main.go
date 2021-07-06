package main

import (
	"flag"
	"os"

	redisql "github.com/DGKSK8LIFE/redisql"
)

var config redisql.Config

func init() {
	copyFlag := flag.NewFlagSet("copy", flag.ExitOnError)
	user := copyFlag.String("user", "root", "MySQL user")
	password := copyFlag.String("password", "", "MySQL password")
	database := copyFlag.String("database", "", "MySQL database")
	table := copyFlag.String("table", "", "MySQL table")
	redisAddr := copyFlag.String("redisaddr", "", "Redis address")
	redisPass := copyFlag.String("redispass", "", "Redis password")
	copyFlag.Parse(os.Args[2:])
	config = redisql.Config{
		SQLUser:     *user,
		SQLPassword: *password,
		SQLDatabase: *database,
		SQLTable:    *table,
		RedisAddr:   *redisAddr,
		RedisPass:   *redisPass,
	}
}

func main() {
	err := config.Copy()
	if err != nil {
		panic(err)
	}
}
