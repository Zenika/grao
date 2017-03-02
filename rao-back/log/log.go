package log

import "log"

/*
 * WIP logging utilities
 */
type ErrorLevel uint8

const (
    FATAL = iota + 1
    ERROR
)

func log(err error, level ErrorLevel) {
  if err != nil {
      handle(err, level)
  }
}

func handle(err error, level ErrorLevel){
  switch level {
  case FATAL:
    log.Fatal("FATAL " + err)
  case ERROR:
    log.Println("ERROR " + err)
  case DEBUG:
    log.Println("DEBUG " + err)
  }
}
