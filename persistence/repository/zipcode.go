package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// ZipCodeRepository ...
type ZipCodeRepository interface {
	GetZipCode(zipCode string) (zipcode domain.ZipCode, err error)
	InsertZipCode(zipCode domain.ZipCode) error
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

func (c *zipCodeRepository) InsertZipCode(zipCode domain.ZipCode) error {

	return c.db.Create(&zipCode).Error
}

func (c *zipCodeRepository) GetZipCode(zipCode string) (zipcode domain.ZipCode, err error) {

	err = c.db.Where("zipcode = ?", zipCode).First(&zipcode).Error

	return
}
