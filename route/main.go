package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xusenlin/notes/controllers/category"
	"net/http"
)


func InitRoute() *gin.Engine {
	router := gin.Default()

	//router.Use(middleware.CORS())


	router.GET("/api/category/", category.Index)
	router.GET("/api/category/show", category.Show)
	router.GET("/api/category/destroy", category.Destroy)
	router.POST("/api/category/store", category.Store)


	router.StaticFile("/", "./public/dist/index.html")
	router.StaticFile("/favicon.ico", "./public/dist/favicon.ico")
	router.StaticFS("/css", http.Dir("./public/dist/css"))
	router.StaticFS("/js", http.Dir("./public/dist/js"))
	router.StaticFS("/img", http.Dir("./public/dist/img"))

	return router
}