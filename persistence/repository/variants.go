package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// VariantsRepository ...
type VariantsRepository interface {
	InsertVariant(size string)
	GetAVariantFromSize(size string) (variant domain.Variants, err error)
}

type variantsRepository struct {
	db *gorm.DB
}

//NewVariantsRepository ...
func NewVariantsRepository(newDB *gorm.DB) VariantsRepository {
	return &variantsRepository{
		db: newDB,
	}
}

func (v *variantsRepository) InsertVariant(size string) {
	variant := domain.Variants{
		Size: size,
	}
	v.db.Create(&variant)
}

func (v *variantsRepository) GetAVariantFromSize(size string) (variant domain.Variants, err error) {
	err = v.db.Where("size = ?", size).First(&variant).Error

	return
}
