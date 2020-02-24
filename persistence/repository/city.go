package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// CityRepository ...
type CityRepository interface {
	InsertCity(city domain.City)
	GetACityByCityCode(cityCode string) (city domain.City, err error)
}

type cityRepository struct {
	db *gorm.DB
}

//NewCityRepository ...
func NewCityRepository(newDB *gorm.DB, provinceRepo ProvinceRepository) CityRepository {
	return &cityRepository{
		db: newDB,
	}
}

func (c *cityRepository) InsertCity(city domain.City) {
	c.db.Create(&city)
}

func (c *cityRepository) GetACityByCityCode(cityCode string) (city domain.City, err error) {
	err = c.db.Where("city_code = ?", cityCode).First(&city).Error
	return
}
