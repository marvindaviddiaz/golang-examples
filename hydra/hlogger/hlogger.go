package hlogger

import (
	"log"
	"os"
	"sync"
)

type hydraLogger struct {
	*log.Logger
	filename string
}

var hlooger *hydraLogger
var once sync.Once

func createLogger(fname string) *hydraLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 777)

	return &hydraLogger{
		filename:fname,
		Logger: log.New(file, "Hydra ", log.LstdFlags),
	}
}

func GetInstance() *hydraLogger {
	once.Do(func() { // This code will execute once on the program
		hlooger = createLogger("hydralogger.log")
	})
	return hlooger
}
