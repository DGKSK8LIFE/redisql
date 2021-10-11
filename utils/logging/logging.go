package logging

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

var logLevelIsSet bool = false
var logFileIsSet bool = false
var loggingLevel uint8 = 0

func SetLogFile(filepath string) {
	if logFileIsSet { 
		panic("Cannot change logfile midway!")
	}

	f, err := os.OpenFile(filepath, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	logFileIsSet = true
	log.SetOutput(f)
}

func InitLogging(level uint8) { 
	if logLevelIsSet {
		panic("Logging level has already been set! It cannot be changed midway.")
	} 

	logLevelIsSet = true
	loggingLevel = level
}


func Log(s string, level uint8) {
	if loggingLevel > 0 && level <= loggingLevel { 
		log.Printf("Verbosity:%d | %s\n", level, s)
	}
}

func LogResultList(results []redis.Cmder, level uint8) {
	if loggingLevel <= 0 || level > loggingLevel { 
		return
	}
	var b strings.Builder

	for _, r := range results {
		fmt.Fprintf(&b, "%s\n", r)
	}
	log.Printf("Verbosity:%d | CHUNK INSERT:\n", level)
	log.Print(b.String())
}


