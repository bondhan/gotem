package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// ZipCodeRepository ...
type ZipCodeRepository interface {
	InsertZipCode(zipCode domain.ZipCode)
}

type zipCodeRepository struct {
	cityRepo CityRepository
	db       *gorm.DB
}

//NewZipCodeRepository ...
func NewZipCodeRepository(newDB *gorm.DB, cityRepo CityRepository) ZipCodeRepository {
	return &zipCodeRepository{
		db: newDB,
	}
}

func (c *zipCodeRepository) InsertZipCode(zipCode domain.ZipCode) {

	c.db.Create(&zipCode)
}
