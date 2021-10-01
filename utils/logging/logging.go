package logging

import (
	"log"
	"os"
)

type Level int32

// Logging level. One means show only a little. Three means show all levels.
const (
	One   = 1
	Two   = 2
	Three = 3
)

var logLevelIsSet bool = false
var logFileIsSet bool = false
var loggingLevel Level

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

func InitLogging(level Level) { 
	if logLevelIsSet {
		panic("Logging level has already been set! It cannot be changed midway.")
	} else if level < 1 || level > 3 {
		panic("Logging level must be between 1 and 3!")
	}

	logLevelIsSet = true
	loggingLevel = level
}


func Log(s string, level Level) {
	if level < 1 || level > 3 {
		panic("Logging level must be between 1 and 3!")
	} else if level <= loggingLevel { 
		log.Printf("Verbosity:%d | %s\n", level, s)
	}
}




