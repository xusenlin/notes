package category

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context)  {
	name := c.DefaultQuery("name","66666")

	c.JSON(200, gin.H{
		"message": name,
	})
}

func Show(c *gin.Context)  {
	name := c.DefaultQuery("name","66666")
	c.JSON(200, gin.H{
		"message": name,
	})
}
func Destroy(c *gin.Context)  {
	name := c.DefaultQuery("name","66666")
	c.JSON(200, gin.H{
		"message": name,
	})
}
func Store(c *gin.Context)  {
	name := c.DefaultQuery("name","66666")
	c.JSON(200, gin.H{
		"message": name,
	})
}