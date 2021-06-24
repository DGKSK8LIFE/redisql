package main

import (
	"flag"
	"fmt"
	"os"
)

/*
Flag format:
	go run main.go migrate --sql=localhost:3306 --redis=6379
*/

var sql string
var redis string

func init() {
	if os.Args[0] == "migrate" {
		flag.StringVar(&sql, "sql", "localhost:3306", "sql port #")
		flag.StringVar(&redis, "redis", "localhost:6379", "redis port #")
		flag.Parse()
	}
}

func main() {
	fmt.Println(sql, redis)
}
