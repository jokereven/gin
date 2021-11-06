package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "shop/index",
		})
	})
	r.GET("/vedio/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "vedio/index",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"not": "found",
		})
	})
	r.Run(":8080")
}
