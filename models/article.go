package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `binging:"required"`
	Content string `binding:"required"`
	Preview string `binding:"required"`
	Likes   int    `gorm:"default:0"`
}
