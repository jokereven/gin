package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// Todo 定义模型
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:root@(localhost)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	// 定义一个全局的DB对象
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	r := gin.Default()

	// 创建数据库
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	// 绑定模型
	DB.AutoMigrate(&Todo{})

	// 进行数据的增删改查

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"massage": "ok",
		})
	})

	v1Group := r.Group("v1")
	{
		// 代办事项

		// 添加事项
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写代办事件 提交 会发放请求到这里
			// 1. 从请求中把数据那出来
			var todo Todo
			c.BindJSON(&todo)
			// 2. 存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
			// 3. 返回相应
		})

		// 查看事项

		// 查看所有的事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// select
			var todo []Todo
			if err = DB.Find(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		// 查看莫一个代办的事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {})

		// 修该事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, todo)
				return
			}
		})

		// 删除事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"massage": "delete",
				})
				return
			}
		})
	}
	r.Run(":9000")
}
