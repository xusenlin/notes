package route

import (
	"github.com/xusenlin/notes/controllers"
	"net/http"
)



func Run()  {
	http.HandleFunc("/api/category", controllers.CategoryTree)



	http.Handle("/", http.FileServer(http.Dir("./public/")))
}