package log

import(
  "log"
  "fmt"
)

/*
 * WIP logging utilities
 */
type ErrorLevel uint8

const (
    FATAL = iota + 1
    ERROR
    WARNING
)

func Error(err error, level ErrorLevel) {
  if err != nil {
      handleError(err, level)
  }
}

func Debug(message string) {
  log.Println(fmt.Sprintf("DEBUG %v", message))
}

func handleError(err error, level ErrorLevel){
  switch level {
  case FATAL:
    log.Fatal(fmt.Sprintf("FATAL %v", err))
  case ERROR:
    log.Println(fmt.Sprintf("ERROR %v", err))
  case WARNING:
    log.Println(fmt.Sprintf("WARNING %v", err))
  }
}
