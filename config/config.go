package config

import (
	"os"

	"github.com/bondhan/gotem/internal/driver"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// NewDbConfig ...
func NewDbConfig() *gorm.DB {
	//init postgresql database
	postgre := driver.NewDbDriver(os.Getenv("POSGRE_DSN"), "postgres")
	db, err := postgre.ConnectDatabase()

	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	return db
}
