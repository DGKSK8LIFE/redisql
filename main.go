package main

import (
	"flag"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

/*
Current CLI flag options:
	go run main.go -sql 3306 -redis 6379
*/

var sqlPort string
var redisPort string

func init() {
	flag.StringVar(&sqlPort, "sql", "localhost:3306", "sql port")
	flag.StringVar(&redisPort, "redis", "localhost:6379", "redis port")
	flag.Parse()
}

func main() {
	redis.NewClient(&redis.Options{
		Addr:     redisPort,
		Password: "",
		DB:       0,
	})
}
