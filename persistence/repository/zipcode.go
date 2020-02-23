package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// ZipCodeRepository ...
type ZipCodeRepository interface {
	InsertZipCode(zipCode string, cityCode string)
}

type zipCodeRepository struct {
	cityRepo CityRepository
	db       *gorm.DB
}

//NewZipCodeRepository ...
func NewZipCodeRepository(newDB *gorm.DB, cityRepo CityRepository) ZipCodeRepository {
	return &zipCodeRepository{
		cityRepo: cityRepo,
		db:       newDB,
	}
}

func (c *zipCodeRepository) InsertZipCode(zipCode string, cityCode string) {
	city, err := c.cityRepo.GetACityByCityCode(cityCode)
	if err != nil {
		log.Error(err)
		return
	}

	zipcode := domain.ZipCode{ZipCode: zipCode, CityID: city.ID}
	c.db.Create(&zipcode)
}
