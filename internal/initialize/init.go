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

	//init mysql database
	// mysql := driver.NewDbDriver(os.Getenv("MYSQL_DSN"), "mysql")
	// db, err := mysql.ConnectDatabase()

	// if err != nil {
	// 	fmt.Println(os.Getenv("MYSQL_DSN"))
	// 	log.Error(err)
	// 	os.Exit(-1)
	// }

	// _ = db

	//init postgresql database
	postgreDsn := "host=" + os.Getenv("host") + " port=" + os.Getenv("port") + " user=" + os.Getenv("user") + " dbname=" + os.Getenv("dbname") + " password=" + os.Getenv("password") + " sslmode=" + os.Getenv("sslmode")
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
