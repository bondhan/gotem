package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// ProvinceRepository ...
type ProvinceRepository interface {
	InsertProvince(province domain.Province)
	GetAProvinceByProvinceCode(provinceCode string) (province domain.Province, err error)
}

type provinceRepository struct {
	countryRepo CountryRepository
	db          *gorm.DB
}

//NewProvinceRepository ...
func NewProvinceRepository(newDB *gorm.DB) ProvinceRepository {
	return &provinceRepository{
		db: newDB,
	}
}

func (p *provinceRepository) InsertProvince(province domain.Province) {
	p.db.Create(&province)
}

func (p *provinceRepository) GetAProvinceByProvinceCode(provinceCode string) (province domain.Province, err error) {
	err = p.db.Where("province_code = ?", provinceCode).First(&province).Error
	return
}
