package domain

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// Seller ...
type Seller struct {
	gorm.Model
	Name      string `gorm:"column:name" json:"name"`
	Mobile    string `gorm:"column:mobile" json:"mobile"`
	Email     string `gorm:"column:email" json:"email"`
	Address   string `gorm:"column:address" json:"address"`
	ZipCodeID uint   `gorm:"column:zipcode_id" json:"zipcode_id"`
}

// TableName sets the insert table name for this struct type
func (m *Seller) TableName() string {
	return "m_seller"
}
