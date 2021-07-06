package main

import (
	"flag"
	"os"

	redisql "github.com/DGKSK8LIFE/redisql"
)

var config redisql.Config

func init() {
	migrateFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	user := migrateFlag.String("user", "root", "MySQL user")
	password := migrateFlag.String("password", "", "MySQL password")
	database := migrateFlag.String("database", "", "MySQL database")
	table := migrateFlag.String("table", "", "MySQL table")
	redisAddr := migrateFlag.String("redisaddr", "", "Redis address")
	redisPass := migrateFlag.String("redispass", "", "Redis password")
	migrateFlag.Parse(os.Args[2:])
	config = redisql.Config{
		User:      *user,
		Password:  *password,
		Database:  *database,
		Table:     *table,
		RedisAddr: *redisAddr,
		RedisPass: *redisPass,
	}
}

func main() {
	err := config.Migrate()
	if err != nil {
		panic(err)
	}
}
