package models

import (
	"github.com/jinzhu/gorm"
	"github.com/xusenlin/notes/database"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"size:50"`
	ParentId    uint
	Order       uint
	Slug        string `gorm:"size:50;unique"`
	Description string
}

func GetAllCategory() []Category {

	var categories []Category

	database.DB.Find(&categories)
	return categories

}

func GetCategoryById(id string) Category {

	var category Category

	database.DB.Where("id = ?",id).First(&category)
	return category

}

func DeleteCategoryById(id string) string {

	var category Category

	database.DB.Where("id = ?",id).Delete(&category)
	return  id
}

//func CreateCategory(c *gin.Context) string {
//
//
//}