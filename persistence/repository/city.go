package repository

import (
	"github.com/jinzhu/gorm"
)

// CityRepository ...
type CityRepository interface {
	InsertCity(CityCode string, CityName string)
}

type cityRepository struct {
	db *gorm.DB
}

//NewCityRepository ...
func NewCityRepository(newDB *gorm.DB) CityRepository {
	return &cityRepository{
		db: newDB,
	}
}

func (c *cityRepository) InsertCity(CityCode string, CityName string) {
}
