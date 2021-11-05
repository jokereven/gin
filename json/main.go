package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 方法一使用map
		data := map[string]interface{}{
			"name": "Joker",
			"age":  19,
		}
		c.JSON(http.StatusOK, data)
	})

	// 方法二使用结构体
	r.GET("/str", func(c *gin.Context) {
		type S struct {
			Name string `json:"name"`
			age  int
		}
		data := S{
			Name: "joker",
			age:  19,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8000")
}
