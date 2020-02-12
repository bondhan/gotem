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

// City ...
type City struct {
	gorm.Model
	CityName   string `gorm:"column:city_name" json:"city_name"`
	ProvinceID uint   `gorm:"column:province_id" json:"province_id"`
}

// TableName sets the insert table name for this struct type
func (m *City) TableName() string {
	return "m_city"
}
