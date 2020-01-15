package repository

import (
	"github.com/bondhan/gotem/models/db/dbmodel"
	"github.com/jinzhu/gorm"
)

// BuyerRepository ...
type BuyerRepository interface {
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

func (b *buyerRepository) GetABuyer(id uint) (dbmodel.Buyer, error) {
	var buyer dbmodel.Buyer
	err := b.db.Where("id = ?", id).First(&buyer).Error

	return buyer, err
}
