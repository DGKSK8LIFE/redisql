package main

import (
	"flag"
	"os"
	"redis-sql/utils/migration"
)

var user *string
var password *string
var database *string
var table *string

func init() {
	migrateFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	user = migrateFlag.String("user", "root", "MySQL user")
	password = migrateFlag.String("password", "", "MySQL password")
	database = migrateFlag.String("database", "", "MySQL database")
	table = migrateFlag.String("table", "", "MySQL table")
	migrateFlag.Parse(os.Args[2:])
}

func main() {
	err := migration.Migrate(*user, *password, *database, *table)
	if err != nil {
		panic(err)
	}
}
