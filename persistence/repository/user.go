package repository

import (
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/jinzhu/gorm"
)

// UserRepository ...
type UserRepository interface {
	GetAUser(id uint) (domain.User, error)
	InsertAUser(user domain.User) error
	GetAUserFromMobile(mobile string) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

//NewUserRepository ...
func NewUserRepository(newDB *gorm.DB) UserRepository {
	return &userRepository{
		db: newDB,
	}
}

func (b *userRepository) GetAUser(id uint) (domain.User, error) {
	var user domain.User
	err := b.db.Where("id = ?", id).First(&user).Error

	return user, err
}

func (b *userRepository) GetAUserFromMobile(mobile string) (domain.User, error) {
	var user domain.User
	err := b.db.Where("mobile = ?", mobile).First(&user).Error

	return user, err
}

func (b *userRepository) InsertAUser(user domain.User) error {
	err := b.db.Create(&user).Error

	return err
}
