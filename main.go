package main

import (
	"github.com/xusenlin/notes/database"
	"github.com/xusenlin/notes/route"
	"log"
	"net/http"
)


func main() {

	if err := database.Open(); err != nil{
		panic("failed to connect database")
	}
	defer database.Close()

	route.Run()

	log.Println("Listening...")
	http.ListenAndServe(":8088", nil)
}

