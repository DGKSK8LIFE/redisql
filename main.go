package main

import (
	"flag"
	"fmt"
)

/*
Current CLI flag options:
	go run main.go -sql=3306 -redis=6379
*/

var sql int
var redis int

func init() {
	flag.IntVar(&sql, "sql", 3306, "sql port")
	flag.IntVar(&redis, "redis", 6379, "redis port")
	flag.Parse()
}

func main() {
	fmt.Printf("sql: %d\nredis: %d\n", sql, redis)
}
