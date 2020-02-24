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

// Buyer ...
type Buyer struct {
	gorm.Model
	UserID uint `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (m *Buyer) TableName() string {
	return "m_buyer"
}
