package main

import (
	"flag"
	"os"

	redisql "github.com/DGKSK8LIFE/redisql"
)

var dataType *string
var file *string

func init() {
	copyFlag := flag.NewFlagSet("copy", flag.ExitOnError)
	dataType = copyFlag.String("type", "", "Data type in Redis")
	file = copyFlag.String("config", "", "Path to config file")
	copyFlag.Parse(os.Args[2:])
}

func main() {
	config, err := redisql.NewConfig(*file)
	if err != nil {
		panic(err)
	}

	switch *dataType {
	case "string":
		if err = config.CopyToString(); err != nil {
			panic(err)
		}
	case "list":
		if err = config.CopyToList(); err != nil {
			panic(err)
		}
	case "hash":
		if err = config.CopyToHash(); err != nil {
			panic(err)
		}
	}
}
