package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"size:50"`
	ParentId    uint
	Order       uint
	Slug        string `gorm:"size:50;unique"`
	Description string
}

