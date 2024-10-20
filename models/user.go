package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"` //用户名不重复
	PassWord string
}
