package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"massage": "ok",
		})
	})
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "/user/index",
			})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "/user/login",
			})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "/user/login",
			})
		})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "/shop/index",
			})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "/shop/cart",
			})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "/shop/checkout",
			})
		})
	}
	r.Run(":8080")
}
