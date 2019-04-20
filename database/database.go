package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Open() error {
	var err error
	DB, err = gorm.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}
	// Migrate the schema
	//DB.AutoMigrate(&models.Category{},&models.Notes{})

	return nil
}

func Close()  {
	DB.Close()
}