package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	ClassID  uint
	Name     string
	Teachers []Teacher `gorm:"many2many:student_teachers"`
	IDcards  IDcard
}

type Teacher struct {
	gorm.Model
	Name     string
	Students []Student `gorm:"many2many:student_teachers"`
}

type IDcard struct {
	gorm.Model
	StudentID uint
	Num       int
}

func main() {
	db, err := gorm.Open("mysql", "root:mobaisilent@tcp(122.51.14.13:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	} else {
		fmt.Println("连接数据库成功")
	}
	defer db.Close()

	r := gin.Default()
	// 使用中间件

	r.GET("/student/:ID", func(c *gin.Context) {
		id := c.Param("ID") // 获取参数
		var student Student
		db.Preload("Teachers").Preload("IDcards").First(&student, "id = ?", id)
		c.JSON(200, gin.H{
			"student": student,
		})
	})

	r.GET("/class/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var class Class
		db.Preload("Students").Preload("Students.Teachers").Preload("Students.IDcards").First(&class, "id = ?", id)
		c.JSON(200, gin.H{
			"class": class,
		})
	})

	r.Run(":8888")
}
