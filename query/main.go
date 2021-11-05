package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// name := c.Query("query")
		// name := c.DefaultQuery("query", "Default")
		name, ok := c.GetQuery("query")
		if !ok {
			name = "Default"
		}
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8000")
}
