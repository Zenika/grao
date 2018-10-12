package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
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

const REFERER  = ""

var level string = os.Getenv("GRAO_LOG_LEVEL")

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

}

func Close() {

}

func Error(err error, level ErrorLevel, refererOptional ...string) {
	referer := ""
	if len(refererOptional) > 0 {
		referer = " - FROM : " + refererOptional[0]
	}
	if err != nil {
		handleError(err, level, referer)
	}
}

func Debug(message string, refererOptional ...string) {
	referer := ""
	if len(refererOptional) > 0 {
		referer = " - FROM : " + refererOptional[0]
	}
	if _level() > DEBUG {
		return
	}
	log.Println(fmt.Sprintf("DEBUG%v - MESSAGE: %v",referer, message))
}

func Info(message string, refererOptional ...string) {
	referer := ""
	if len(refererOptional) > 0 {
		referer = " - FROM : " + refererOptional[0]
	}

	if _level() > INFO {
		return
	}
	log.Println(fmt.Sprintf("INFO%v : %v", referer, message))
}

func handleError(err error, level ErrorLevel, referer string) {
	switch level {
	case FATAL:
		log.Fatal(fmt.Sprintf("FATAL%v - TRACE: %v", referer, err))
	case ERROR:
		log.Println(fmt.Sprintf("ERROR%v - TRACE: %v", referer, err))
	case WARNING:
		log.Println(fmt.Sprintf("WARNING%v - TRACE: %v", referer, err))
	}
}

func GetReferer() string {
	_, referer, _, ok := runtime.Caller(1)
	if ok {
		return referer
	}
	return "None found"
}
