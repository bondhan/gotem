package errorhandler

import (
	"os"
	"sync"

	"github.com/bondhan/gotem/internal/driver"
	log "github.com/sirupsen/logrus"
)

var (
	once sync.Once
)

func newLogHandler() {
	//init the log
	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		driver.NewLogDriver(os.Getenv("LOG_NAME"), log.ErrorLevel).InitLog()
	} else {
		driver.NewLogDriver(os.Getenv("LOG_NAME"), log.TraceLevel).InitLog()
	}
}

func DoLog(level log.Level, err error) {
	once.Do(func() {
		newLogHandler()
	})

	if err != nil {
		log.Debug(err)
	}

}
