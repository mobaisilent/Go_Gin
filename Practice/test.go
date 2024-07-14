package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Sex  string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:mobaisilent@tcp(122.51.14.13:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	} else {
		fmt.Println("建立连接成功")
	}
	defer db.Close() // defer关闭
	db.AutoMigrate(&HelloWorld{})

	result := db.Where("id = ?", 1).Delete(&HelloWorld{})
	if result.Error != nil {
		fmt.Println("删除失败:", result.Error)
	} else {
		fmt.Println("删除成功")
	}
}
