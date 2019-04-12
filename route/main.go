package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xusenlin/notes/controllers/category"
	"net/http"
)


func InitRoute() *gin.Engine {
	router := gin.Default()


	categoryGroup := router.Group("/api/category")
	{
		categoryGroup.GET("/", category.Index)
		categoryGroup.GET("/", category.Show)
		categoryGroup.GET("/", category.Destroy)
		categoryGroup.POST("/", category.Store)
	}

	router.StaticFile("/", "./public/dist/index.html")
	router.StaticFile("/favicon.ico", "./public/dist/favicon.ico")
	router.StaticFS("/css", http.Dir("./public/dist/css"))
	router.StaticFS("/js", http.Dir("./public/dist/js"))
	router.StaticFS("/img", http.Dir("./public/dist/img"))

	return router
}