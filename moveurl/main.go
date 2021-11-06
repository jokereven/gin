package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://blog.code520.com.cn/")
	})
	r.GET("/one", func(c *gin.Context) {
		c.Request.URL.Path = "/two"
		r.HandleContext(c)
	})
	r.GET("/two", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.Run(":8000")
}
