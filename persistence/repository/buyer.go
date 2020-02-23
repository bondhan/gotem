package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// BuyerRepository ...
type BuyerRepository interface {
	GetABuyer(id uint) (domain.Buyer, error)
	InsertABuyer(buyer domain.Buyer) error
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

func (b *buyerRepository) InsertABuyer(buyer domain.Buyer) error {
	err := b.db.Create(&buyer).Error

	return err
}
