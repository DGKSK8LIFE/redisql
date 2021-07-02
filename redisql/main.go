package main

import (
	"flag"
	"os"

	redisql "github.com/DGKSK8LIFE/redisql"
)

var user *string
var password *string
var database *string
var table *string
var redisAddr *string
var redisPass *string

func init() {
	migrateFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	user = migrateFlag.String("user", "root", "MySQL user")
	password = migrateFlag.String("password", "", "MySQL password")
	database = migrateFlag.String("database", "", "MySQL database")
	table = migrateFlag.String("table", "", "MySQL table")
	redisAddr = migrateFlag.String("redisaddr", "", "Redis address")
	redisPass = migrateFlag.String("redispass", "", "Redis password")
	migrateFlag.Parse(os.Args[2:])
}

func main() {
	err := redisql.Migrate(*user, *password, *database, *table, *redisAddr, *redisPass)
	if err != nil {
		panic(err)
	}
}
