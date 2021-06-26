package main

import (
	"redis-sql/utils/migration"
)

/*
Current CLI flag options:
	go run main.go -sql 3306 -redis 6379
*/

// var sqlPort string
// var redisPort string

// func init() {
// 	// flag.StringVar(&sqlPort, "sql", "localhost:3306", "sql port")
// 	// flag.StringVar(&redisPort, "redis", "localhost:6379", "redis port")
// 	// flag.Parse()
// }

func main() {
	err := migration.Migrate("root", "celebrities", "celebrity", true)
	if err != nil {
		panic(err)
	}
}
