package log

import (
	"fmt"
	"log"
	"os"
)

/*
 * WIP logging utilities
 */
type ErrorLevel uint8

const (
	DEBUG = iota + 1
	INFO
	WARNING
	ERROR
	FATAL
)

var level string = os.Getenv("GRAO_LOG_LEVEL")
var file os.File

func _level() int {
	switch level {
	case "FATAL":
		return FATAL
	case "ERROR":
		return ERROR
	case "WARNING":
		return WARNING
	case "INFO":
		return INFO
	}
	return DEBUG
}

func Init() {
	dest := os.Getenv("GRAO_LOG_FILE")
	if "" == dest {
		dest = "rao.log"
	}
	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func Close() {
	file.Close()
}

func Error(err error, level ErrorLevel) {
	if err != nil {
		handleError(err, level)
	}
}

func Debug(message string) {
	if _level() > DEBUG {
		return
	}
	log.Println(fmt.Sprintf("DEBUG %v", message))
}

func Info(message string) {
	if _level() > INFO {
		return
	}
	log.Println(fmt.Sprintf("INFO %v", message))
}

func handleError(err error, level ErrorLevel) {
	switch level {
	case FATAL:
		log.Fatal(fmt.Sprintf("FATAL %v", err))
	case ERROR:
		log.Println(fmt.Sprintf("ERROR %v", err))
	case WARNING:
		log.Println(fmt.Sprintf("WARNING %v", err))
	}
}
