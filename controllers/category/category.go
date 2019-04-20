package category

import (
	"github.com/gin-gonic/gin"
	"github.com/xusenlin/notes/models"
	"net/http"
)


func Index(c *gin.Context)  {

	categories := models.GetAllCategory()

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func Show(c *gin.Context)  {

	categoryId := c.DefaultQuery("id","0")
	category := models.GetCategoryById(categoryId)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func Destroy(c *gin.Context)  {
	categoryId := c.DefaultQuery("id","0")

	deleteId := models.DeleteCategoryById(categoryId)

	c.JSON(http.StatusOK, gin.H{"data": deleteId})
}


func Store(c *gin.Context)  {

	var category models.Category

	c.BindJSON(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}