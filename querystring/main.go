package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserInfo 定义结构体
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// 	username := c.Query("username")
		// 	password := c.Query("password")
		// 	u := UserInfo{
		// 		username: username,
		// 		password: password,
		// 	}
		// 	fmt.Println(u)
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"massage": "ok",
		// 	})
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Println(u)
			c.JSON(http.StatusOK, gin.H{
				"massage": "ok",
			})
		}
	})
	r.POST("/json", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Println(u)
			c.JSON(http.StatusOK, gin.H{
				"massage": "ok",
			})
		}
	})
	r.Run(":8000")
}
