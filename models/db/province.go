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

// Province ...
type Province struct {
	gorm.Model
	ProvinceName string `gorm:"column:province_name" json:"province_name"`
	ProvinceCode string `gorm:"column:province_code" json:"province_code"`
	CountryID    uint   `gorm:"column:country_id" json:"country_id"`
}

// TableName sets the insert table name for this struct type
func (m *Province) TableName() string {
	return "m_province"
}
