package logger

import (
	"log"
	"os"
)

var (
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
)

func init() {
	info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	error = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Error(v ... interface{}) {
	error.Println(v)
}

func Errorf(format string, v ... interface{}) {
	error.Printf(format, v)
}
