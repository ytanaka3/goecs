package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", Hello)
	v1 := router.Group("api/v1")
	{
		v1.GET("/hello", Hello)
		v1.GET("/channels", List)
		v1.GET("/channels/:id", Show)
		v1.POST("/channels/new", New)
		v1.PUT("/channels/:id", Update)
		v1.DELETE("/channels/:id", Delete)
	}

	return router
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"ok": "Delete id:" + id})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"ok": "Update" + id})
}

func List(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "Show list"})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"ok": id})
}

func New(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "Created New"})
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "Hello API!"})
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
