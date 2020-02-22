package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// BuyerRepository ...
type BuyerRepository interface {
	GetABuyer(id uint) (domain.Buyer, error)
}

type buyerRepository struct {
	db *gorm.DB
}

//NewBuyerRepository ...
func NewBuyerRepository(newDB *gorm.DB) BuyerRepository {
	return &buyerRepository{
		db: newDB,
	}
}

func (b *buyerRepository) GetABuyer(id uint) (domain.Buyer, error) {
	var buyer domain.Buyer
	err := b.db.Where("id = ?", id).First(&buyer).Error

	return buyer, err
}
