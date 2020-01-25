package initialize

import (
	"fmt"
	"os"

	"github.com/bondhan/gotem/internal/driver"
	log "github.com/sirupsen/logrus"
)

/*Init ... Initialize the application server  */
func Init() {

	//init the log
	driver.NewLogDriver(os.Getenv("LOG_NAME"), log.TraceLevel).InitLog()
	//init postgresql database
	postgreDsn := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=" + os.Getenv("DB_SSLMODE")
	postgre := driver.NewDbDriver(postgreDsn, "postgres")
	db, err := postgre.ConnectDatabase()

	if err != nil {
		fmt.Println(postgreDsn)
		log.Error(err)
		os.Exit(-1)
	}

	_ = db

	//init the repository

	//init rabbit mq

	//init the redis

	//init the server

}
