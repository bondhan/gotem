package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// CityRepository ...
type CityRepository interface {
	InsertCity(cityCode string, cityName string, provinceCode string)
	GetACityByCityCode(cityCode string) (city domain.City, err error)
}

type cityRepository struct {
	provinceRepo ProvinceRepository
	db           *gorm.DB
}

//NewCityRepository ...
func NewCityRepository(newDB *gorm.DB, provinceRepo ProvinceRepository) CityRepository {
	return &cityRepository{
		provinceRepo: provinceRepo,
		db:           newDB,
	}
}

func (c *cityRepository) InsertCity(cityCode string, cityName string, provinceCode string) {
	province, err := c.provinceRepo.GetAProvinceByProvinceCode(provinceCode)
	if err != nil {
		log.Error(err)
		return
	}
	city := domain.City{CityCode: cityCode, CityName: cityName, ProvinceID: province.ID}
	c.db.Create(&city)
}

func (c *cityRepository) GetACityByCityCode(cityCode string) (city domain.City, err error) {
	err = c.db.Where("city_code = ?", cityCode).First(&city).Error
	return
}
