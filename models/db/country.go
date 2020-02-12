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

// Country ...
type Country struct {
	gorm.Model
	CountryName string `gorm:"column:country_name" json:"country_name"`
	CountryCode string `gorm:"column:country_code" json:"country_code"`
}

// TableName sets the insert table name for this struct type
func (m *Country) TableName() string {
	return "m_country"
}
