package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// ProductsRepository ...
type ProductsRepository interface {
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

func (b *productsRepository) GetABuyer(id uint) (domain.Products, error) {

	var products domain.Products

	err := b.db.Where("id = ?", id).First(&products).Error

	return products, err
}
