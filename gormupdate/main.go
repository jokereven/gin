package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义模型
type User struct {
	gorm.Model
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

	// 查询
	var user User
	db.First(&user)
	// 更新
	user.Name = "Kiku"
	user.Age = 32
	db.Debug().Save(&user)
	db.Debug().Model(&user).Update("name", "code-server")
	// updates
	m1 := map[string]interface{}{
		"name": "莫克",
		"age":  18,
	}
	db.Model(&user).Debug().Updates(m1)
	db.Model(&user).Debug().Select("age").Update(m1)
	db.Model(&user).Omit("age").Updates(m1)
	// by sql to update
	db.Model(&User{}).Update("age", gorm.Expr("age + ?", 2))
}
