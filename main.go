package main

import (
	"github.com/xusenlin/notes/database"
	"github.com/xusenlin/notes/route"

)


func main() {

	if err := database.Open(); err != nil{
		panic("failed to connect database")
	}
	defer database.Close()

	router := route.InitRoute()

	router.Run(":9035")
}

