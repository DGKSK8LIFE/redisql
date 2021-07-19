package main

import (
	"flag"
	"io/ioutil"
	"os"

	redisql "github.com/DGKSK8LIFE/redisql"
	"gopkg.in/yaml.v2"
)

var config redisql.Config
var dataType *string
var file *string

func init() {
	copyFlag := flag.NewFlagSet("copy", flag.ExitOnError)
	dataType = copyFlag.String("type", "list", "Data type in Redis")
	file = copyFlag.String("config", "", "Path to config file")
	copyFlag.Parse(os.Args[2:])
}

func main() {
	yamlFile, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
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
