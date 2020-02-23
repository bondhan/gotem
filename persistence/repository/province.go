package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// ProvinceRepository ...
type ProvinceRepository interface {
	InsertProvince(provinceCode string, provinceName string, countryCode string) error
	GetAProvinceByProvinceCode(provinceCode string) (province domain.Province, err error)
}

type provinceRepository struct {
	countryRepo CountryRepository
	db          *gorm.DB
}

//NewProvinceRepository ...
func NewProvinceRepository(newDB *gorm.DB, countryRepo CountryRepository) ProvinceRepository {
	return &provinceRepository{
		countryRepo: countryRepo,
		db:          newDB,
	}
}

func (p *provinceRepository) InsertProvince(provinceCode string, provinceName string, countryCode string) error {
	country, err := p.countryRepo.GetIDByCountryCode(countryCode)
	if err != nil {
		log.Error(err)

		return err
	}

	province := domain.Province{ProvinceCode: provinceCode, ProvinceName: provinceName, CountryID: country.ID}
	p.db.Save(&province)

	return nil
}

func (p *provinceRepository) GetAProvinceByProvinceCode(provinceCode string) (province domain.Province, err error) {
	err = p.db.Where("province_code = ?", provinceCode).First(&province).Error
	return
}
