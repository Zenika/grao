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
	WARNING
	ERROR
	FATAL
)

var level = os.Getenv("RAO_LOG_LEVEL")

func _level() int {
	switch level {
	case "FATAL":
		return FATAL
	case "ERROR":
		return ERROR
	case "WARNING":
		return WARNING
	}
	return DEBUG
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
