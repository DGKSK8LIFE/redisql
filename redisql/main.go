package main

import (
	"flag"
	"io/ioutil"
	"os"

	redisql "github.com/DGKSK8LIFE/redisql"
	"gopkg.in/yaml.v2"
)

var config redisql.Config
var file *string

func init() {
	copyFlag := flag.NewFlagSet("copy", flag.ExitOnError)
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
	if err = config.CopyToHash(); err != nil {
		panic(err)
	}
}
