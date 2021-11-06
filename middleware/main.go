package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func indexHander(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"massage": "ok",
	})
}

func main() {
	r := gin.Default()
	// 路由组 添加中间件
	// userGroup := r.Group("/user")
	// 方法一
	// userGroup := r.Group("/user", StatCost())

	// 方法二
	userGroup := r.Group("/user")
	userGroup.Use(StatCost())
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "ok/index",
			})
		})
		userGroup.GET("/code520", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "ok/code520",
			})
		})
	}
	// 全局注册中间件
	r.Use(StatCost())
	r.GET("/index", indexHander)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"massage": "test"})
	})
	r.GET("/onlyone", StatCost(), func(c *gin.Context) {
		fmt.Println("this is a onlyone router")
		c.JSON(http.StatusOK, gin.H{
			"massage": "this is onlyone router",
		})
	})
	r.Run(":8080")
}
