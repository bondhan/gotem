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

// ZipCode ...
type ZipCode struct {
	gorm.Model
	ZipCode string `gorm:"column:zip_code" json:"zip_code"`
	CityID  uint   `gorm:"column:city_id" json:"city_id"`
}

// TableName sets the insert table name for this struct type
func (m *ZipCode) TableName() string {
	return "m_zipcode"
}
