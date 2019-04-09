package main

import (
	"github.com/xusenlin/notes/database"
	"net/http"
)


func main() {

	if err := database.Open(); err != nil{
		panic("failed to connect database")
	}
	defer database.Close()

	http.ListenAndServe(":8088", http.FileServer(http.Dir("./public/")))

}
