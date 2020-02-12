package dbmodel

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

// Variants ...
type Variants struct {
	gorm.Model
	Size string `gorm:"column:size" json:"size"`
}

// TableName sets the insert table name for this struct type
func (m *Variants) TableName() string {
	return "m_variants"
}
