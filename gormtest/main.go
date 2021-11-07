package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义模型
type User struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	// u1 := User{
	// 	Name: "mooqle",
	// 	Age:  22,
	// }
	// db.Create(&u1)

	// u2 := User{
	// 	Name: "jokereven",
	// 	Age:  19,
	// }
	// db.Create(&u2)

	// 查询第一个
	// var user User
	// db.First(&user)
	// fmt.Println(user)

	// 查询第指定limit的数据
	// var takeuser []User
	// db.Take(&takeuser)
	// fmt.Println(takeuser)

	// 查询所有 select all
	// var finduser []User
	// db.Find(&finduser)
	// fmt.Println(finduser)

	// var firstuser User
	// db.First(&firstuser, 1)
	// fmt.Println(firstuser)

	// 普通sql查询
	var ordinarysql User
	db.Where("name = ?", "jokereven").First(&ordinarysql)
	fmt.Println(ordinarysql)

	//struct and map
	var structmap User
	db.Where(&User{Name: "mooqle"}).First(&structmap)
	fmt.Println(structmap)
}
