package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// CountryRepository ...
type CountryRepository interface {
	InsertCountry(countryCode string, countryName string)
	GetACountryByCountryCode(countryCode string) (country domain.Country, err error)
}

type countryRepository struct {
	db *gorm.DB
}

//NewCountryRepository ...
func NewCountryRepository(newDB *gorm.DB) CountryRepository {
	return &countryRepository{
		db: newDB,
	}
}

func (c *countryRepository) InsertCountry(countryCode string, countryName string) {
	country := domain.Country{CountryName: countryName, CountryCode: countryCode}
	c.db.Create(&country)
}

func (c *countryRepository) GetACountryByCountryCode(countryCode string) (country domain.Country, err error) {
	err = c.db.Where("country_code = ?", countryCode).First(&country).Error
	return
}
