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
	//
	db.AutoMigrate(&User{})

	// create
	// u := User{
	// 	Name: "Jason",
	// 	Age:  24,
	// }
	// fmt.Println(db.NewRecord(u))
	// db.Create(&u)
	// fmt.Println(db.NewRecord(u))

	// 查询

	// var user User
	// db.First(&user)
	// fmt.Println(user)

	var user []User
	db.Take(&user)
	fmt.Println(user)

	// var user []User
	// db.Find(&user)
	// fmt.Println(user)



}0
