package main

import (
	"flag"
	"fmt"
)

/*
Flag format:
	go run main.go migrate --sql=localhost:3306 --redis=6379
*/

var sql int
var redis int

func init() {
	flag.IntVar(&sql, "sql", 3306, "sql port")
	flag.IntVar(&redis, "redis", 6379, "redis port")
	flag.Parse()
}

func main() {
	fmt.Println(sql, redis)
}