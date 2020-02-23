package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// SellerRepository ...
type SellerRepository interface {
	GetASeller(id uint) (domain.Seller, error)
	InsertASeller(Seller domain.Seller) error
}

type sellerRepository struct {
	db *gorm.DB
}

//NewSellerRepository ...
func NewSellerRepository(newDB *gorm.DB) SellerRepository {
	return &sellerRepository{
		db: newDB,
	}
}

func (b *sellerRepository) GetASeller(id uint) (domain.Seller, error) {
	var Seller domain.Seller
	err := b.db.Where("id = ?", id).First(&Seller).Error

	return Seller, err
}

func (b *sellerRepository) InsertASeller(Seller domain.Seller) error {
	err := b.db.Create(&Seller).Error

	return err
}
