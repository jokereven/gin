package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func m(c *gin.Context) {
	fmt.Println("this is a middleware")
	start := time.Now()
	c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	// 调用该请求的剩余处理程序
	c.Next()
	// 不调用该请求的剩余处理程序
	// c.Abort()
	// 计算耗时
	cost := time.Since(start)
	log.Println(cost)
}

func indexHander(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"massage": "ok",
	})
}

func main() {
	r := gin.Default()
	// 全局注册中间件
	r.Use(m)
	r.GET("/index", indexHander)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"massage": "test",
		})
	})
	r.Run(":8080")
}
