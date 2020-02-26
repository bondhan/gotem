package errorhandler

import (
	"os"

	"github.com/bondhan/gotem/internal/driver"
	log "github.com/sirupsen/logrus"
)

// NewLogHandler creating log handler
func NewLogHandler() error {
	//init the log
	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		driver.NewLogDriver(os.Getenv("LOG_NAME"), log.ErrorLevel).InitLog()
	} else {
		driver.NewLogDriver(os.Getenv("LOG_NAME"), log.TraceLevel).InitLog()
	}

	return nil
}

//
func DoLog(level log.Level, err error) {
	if err != nil {
		log.Debug(err)
	}
}
