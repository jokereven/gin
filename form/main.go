package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// post请求
	r.POST("/login", func(c *gin.Context) {
		// username := c.PostForm("username")
		// password := c.PostForm("password")
		// username := c.DefaultPostForm("username", "default")
		// password := c.DefaultPostForm("passwd", "******")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "无效用户名"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "无效的密码"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})
	r.Run(":8000")
}
