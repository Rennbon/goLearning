package mysql

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Gender   bool
	Email    string `gorm:"type:varchar(100);unique_index"`
	Address  string `gorm:"size:255"`
	Password string `gorm:"size:20"`
}
