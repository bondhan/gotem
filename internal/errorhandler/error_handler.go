package errorhandler

import (
	"os"

	"github.com/bondhan/gotem/internal/driver"
	"github.com/sirupsen/logrus"
)

var (
	log *driver.LogDriver
)

// NewLogHandler creating log handler
func NewLogHandler() error {
	//init the log
	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		log = driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.ErrorLevel)

	} else {
		log = driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.TraceLevel)

	}
	log.InitLog()

	return nil
}

func DoLog(level logrus.Level, err error) {
	switch level {
	}

}
