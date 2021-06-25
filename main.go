package main

import (
	"flag"
)

/*
Current CLI flag options:
	go run main.go -sql 3306 -redis 6379
*/

var sqlPort int
var redisPort int

func init() {
	flag.IntVar(&sqlPort, "sql", 3306, "sql port")
	flag.IntVar(&redisPort, "redis", 6379, "redis port")
	flag.Parse()
}

func main() {
}
