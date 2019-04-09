package route

import (
	"github.com/xusenlin/notes/controller"
	"net/http"
)



func Run()  {
	http.HandleFunc("/api/category", controller.CategoryTree)



	http.Handle("/", http.FileServer(http.Dir("./public/")))
}