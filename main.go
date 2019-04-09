package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price int
}

func main() {
	db, err := gorm.Open("sqlite3", "./database/database.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 创建
	//for i:= 1;i <101;i++ {
	//	db.Create(&Product{Code: "xusenlin", Price: i})
	//}


	//读取
	//var product Product
	//db.First(&product, 1) // 查询id为1的product
	//
	//fmt.Println(product)
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//
	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删除 - 删除product
	//db.Delete(&product)
}