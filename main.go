package main

import "github.com/gin-gonic/gin"

// Gin 框架的使用

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "Hello World",
	})
}

func main() {
	r := gin.Default() //返回默认的路由引擎

	// 指定用户使用Get请求访问/hello，执行sayHello这个函数
	r.GET("/hello", sayHello)

	// Start Serve
	// r.Run()

	// Start On Assign Port
	r.Run(":8000")
}
