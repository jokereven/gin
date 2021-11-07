package main

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义模型
type User struct {
	ID int64
	// Name *string `gorm:"default:'小王子'"`
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int16
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	u := User{
		// Name: new(string),
		Name: sql.NullString{String: "", Valid: true},
		Age:  23,
	}
	db.Debug().Create(&u)
}
