package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义模型
type User struct {
	ID   int64
	Name string
	Age  int
	gorm.Model
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

	// 删除
	// var u = User{}
	// u.ID = 2
	// db.Debug().Delete(&u)
}
