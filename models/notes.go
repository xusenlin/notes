package models

import "github.com/jinzhu/gorm"

type Notes struct {
	gorm.Model
	Title      string
	CategoryId uint
	Order      uint
	Body       string `gorm:"type:TEXT"`
}
