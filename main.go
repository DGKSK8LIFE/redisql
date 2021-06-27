package main

import (
	"redis-sql/utils/migration"
)

func main() {
	err := migration.Migrate("root", "", "celebrities", "celebrity")
	if err != nil {
		panic(err)
	}
}
