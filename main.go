package main

import (
	"flag"
	"fmt"
)

/*
Flag format:
	go run main.go -migrate --sql=localhost:3306 --redis=6379 --options={"uuid": true, "ttl": 120}
*/

func init() {
	migrate := flag.Bool("migrate", false, "migrate from mysql to redis")
	flag.Parse()

	if *migrate {
		fmt.Println(*migrate)
	}
}

func main() {

}
