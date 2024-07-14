package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		file := form.File["file"]
		fmt.Println(file) // 打印二进制信息
		for _, f := range file {
			fmt.Println(f.Filename)
		}
		// 循环打印文件名称
	})
	r.Run(":1234")
}
