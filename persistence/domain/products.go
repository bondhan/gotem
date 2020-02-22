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

// Products ...
type Products struct {
	gorm.Model
	Name        string  `gorm:"column:name" json:"name"`
	Category    string  `gorm:"column:category" json:"category"`
	VariantID   uint    `gorm:"column:variant_id" json:"variant_id"`
	Price       float32 `gorm:"column:price" json:"price"`
	Description string  `gorm:"column:description" json:"description"`
	Quantity    int32   `gorm:"column:quantity" json:"quantity"`
}

// TableName sets the insert table name for this struct type
func (m *Products) TableName() string {
	return "m_products"
}
