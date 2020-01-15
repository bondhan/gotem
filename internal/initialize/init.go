package initialize

import (
	"os"

	"github.com/bondhan/gotem/internal/driver"
	log "github.com/sirupsen/logrus"
)

/*Init ... Initialize the application server  */
func Init() {

	//init the log
	driver.NewLogDriver(os.Getenv("LOG_NAME"), log.TraceLevel).InitLog()

	//init database
	mysql := driver.NewDbDriver(os.Getenv("MYSQL_DSN"), "mysql")
	db, err := mysql.ConnectDatabase()

	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	_ = db

	//init the repository

	//init rabbit mq

	//init the redis

	//init the server

}
