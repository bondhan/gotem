package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// ProductsRepository ...
type ProductsRepository interface {
	AddAproduct(product domain.Products) error
}

type productsRepository struct {
	db *gorm.DB
}

//NewProductsRepository ...
func NewProductsRepository(newDB *gorm.DB) ProductsRepository {
	return &productsRepository{
		db: newDB,
	}
}

func (d *productsRepository) AddAproduct(product domain.Products) error {
	err := d.db.Create(&product).Error
	return err
}
