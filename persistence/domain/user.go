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

// User ...
type User struct {
	gorm.Model
	Name     string `gorm:"column:name" json:"name"`
	Mobile   string `gorm:"column:mobile" json:"mobile"`
	Email    string `gorm:"column:email" json:"email"`
	Address  string `gorm:"column:address" json:"address"`
	ZipCode  string `gorm:"column:zipcode" json:"zipcode"`
	Password string `gorm:"column:password" json:"password"`
	Salt     string `gorm:"column:salt" json:"salt"`
}

// TableName sets the insert table name for this struct type
func (m *User) TableName() string {
	return "m_user"
}
