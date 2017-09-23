package logger

import (
	"github.com/apsdehal/go-logger"
	"os"
)

var log *logger.Logger = nil

func GetInstance() *logger.Logger {
	if (log != nil) {
		return log
	} else {
		var err error
		log, err = logger.New("Airttp", 1, os.Stdout)
		if err != nil {
			panic(err) // Check for error
		}
		return log
	}
}