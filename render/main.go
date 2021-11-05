package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLFiles("./templates/index.tmpl")
	// 解析静态文件
	// 在.tmpl文件中找到以/static就指向./static
	r.Static("/static", "./static")
	// gin框架中添加自定义模板
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "freelancer",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://code520.com.cn'>JokerEven</a>",
		})
	})
	r.Run(":8000")
}
