package logging

import (
	"log"
	"os"
)

var logLevelIsSet bool = false
var logFileIsSet bool = false
var loggingLevel uint32 = 0

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

func InitLogging(level uint32) { 
	if logLevelIsSet {
		panic("Logging level has already been set! It cannot be changed midway.")
	} 

	logLevelIsSet = true
	loggingLevel = level
}


func Log(s string, level uint32) {
	if loggingLevel > 0 && level <= loggingLevel { 
		log.Printf("Verbosity:%d | %s\n", level, s)
	}
}




