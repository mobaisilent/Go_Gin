## 	Gin简介

### Gin 是什么？

Gin 是一个用 Go (Golang) 编写的 HTTP web 框架。 它是一个类似于 martini 但拥有更好性能的 API 框架，由于 httprouter，速度提高了近 40 倍。如果你需要极好的性能，使用 Gin 吧。

### 如何使用 Gin？

下面罗列了一些众所周知的 Gin 使用者.(感兴趣的可以去一一看看)

- [AsciiJSON](https://gin-gonic.com/zh-cn/docs/examples/ascii-json/)
- [HTML 渲染](https://gin-gonic.com/zh-cn/docs/examples/html-rendering/)
- [HTTP2 server 推送](https://gin-gonic.com/zh-cn/docs/examples/http2-server-push/)
- [JSONP](https://gin-gonic.com/zh-cn/docs/examples/jsonp/)
- [Multipart/Urlencoded 绑定](https://gin-gonic.com/zh-cn/docs/examples/multipart-urlencoded-binding/)
- [Multipart/Urlencoded 表单](https://gin-gonic.com/zh-cn/docs/examples/multipart-urlencoded-form/)
- [PureJSON](https://gin-gonic.com/zh-cn/docs/examples/pure-json/)
- [Query 和 post form](https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/)
- [SecureJSON](https://gin-gonic.com/zh-cn/docs/examples/secure-json/)
- [XML/JSON/YAML/ProtoBuf 渲染](https://gin-gonic.com/zh-cn/docs/examples/rendering/)
- [安全页眉](https://gin-gonic.com/zh-cn/docs/examples/security-headers/)
- [绑定 HTML 复选框](https://gin-gonic.com/zh-cn/docs/examples/bind-html-checkbox/)
- [绑定 Uri](https://gin-gonic.com/zh-cn/docs/examples/bind-uri/)
- [绑定表单数据至自定义结构体](https://gin-gonic.com/zh-cn/docs/examples/bind-form-data-request-with-custom-struct/)
- [绑定查询字符串或表单数据](https://gin-gonic.com/zh-cn/docs/examples/bind-query-or-post/)
- [不使用默认的中间件](https://gin-gonic.com/zh-cn/docs/examples/without-middleware/)
- [查询字符串参数](https://gin-gonic.com/zh-cn/docs/examples/querystring-param/)
- [从 reader 读取数据](https://gin-gonic.com/zh-cn/docs/examples/serving-data-from-reader/)
- [定义路由日志的格式](https://gin-gonic.com/zh-cn/docs/examples/define-format-for-the-log-of-routes/)
- [多模板](https://gin-gonic.com/zh-cn/docs/examples/multiple-template/)
- [将 request body 绑定到不同的结构体中](https://gin-gonic.com/zh-cn/docs/examples/bind-body-into-dirrerent-structs/)
- [静态文件服务](https://gin-gonic.com/zh-cn/docs/examples/serving-static-files/)
- [静态资源嵌入](https://gin-gonic.com/zh-cn/docs/examples/bind-single-binary-with-template/)
- [控制日志输出颜色](https://gin-gonic.com/zh-cn/docs/examples/controlling-log-output-coloring/)
- [路由参数](https://gin-gonic.com/zh-cn/docs/examples/param-in-path/)
- [路由组](https://gin-gonic.com/zh-cn/docs/examples/grouping-routes/)
- [模型绑定和验证](https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/)
- [如何记录日志](https://gin-gonic.com/zh-cn/docs/examples/write-log/)
- [上传文件](https://gin-gonic.com/zh-cn/docs/examples/upload-file/)
- [设置和获取 Cookie](https://gin-gonic.com/zh-cn/docs/examples/cookie/)
- [使用 BasicAuth 中间件](https://gin-gonic.com/zh-cn/docs/examples/using-basicauth-middleware/)
- [使用 HTTP 方法](https://gin-gonic.com/zh-cn/docs/examples/http-method/)
- [使用中间件](https://gin-gonic.com/zh-cn/docs/examples/using-middleware/)
- [映射查询字符串或表单参数](https://gin-gonic.com/zh-cn/docs/examples/map-as-querystring-or-postform/)
- [优雅地重启或停止](https://gin-gonic.com/zh-cn/docs/examples/graceful-restart-or-stop/)
- [运行多个服务](https://gin-gonic.com/zh-cn/docs/examples/run-multiple-service/)
- [在中间件中使用 Goroutine](https://gin-gonic.com/zh-cn/docs/examples/goroutines-inside-a-middleware/)
- [支持 Let's Encrypt](https://gin-gonic.com/zh-cn/docs/examples/support-lets-encrypt/)
- [只绑定 url 查询字符串](https://gin-gonic.com/zh-cn/docs/examples/only-bind-query-string/)
- [重定向](https://gin-gonic.com/zh-cn/docs/examples/redirects/)
- [自定义 HTTP 配置](https://gin-gonic.com/zh-cn/docs/examples/custom-http-config/)
- [自定义日志文件](https://gin-gonic.com/zh-cn/docs/examples/custom-log-format/)
- [自定义验证器](https://gin-gonic.com/zh-cn/docs/examples/custom-validators/)
- [自定义中间件](https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/)

## Hello_Gin

```go
package main

import "github.com/gin-gonic/gin"

// 导入gin包

func main() {
	r := gin.Default() // 创建一个默认的路由引擎（创建中间件）
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! This is a Gin server.",
		})
	})
	// 创建一个GET路由，访问/hello时，返回json数据
	r.Run(":28081")
	// 运行服务不加:28081则默认在8080
}

// 使用gin很简单就启动了一个http服务
```

> 很轻松就创建了一个go的HTTP服务，当127.0.0.1:28081后面接入/hello的时候就会返回一段JSON文本，如下所示
>
> 当然就单纯访问 127.0.0.1:28081 会出现404 page not found  ：： 404也说明了后端服务已经启动好了

![image-20240713153525454](./images/image-20240713153525454.png)

## 四种常见请求

示例：

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Handle GET request at /someGet
    router.GET("/someGet", func(c *gin.Context) {
        c.String(200, "GET")
    })

    // Handle POST request at /somePost
    router.POST("/somePost", func(c *gin.Context) {
        c.String(200, "POST")
    })

    // Handle PUT request at /somePut
    router.PUT("/somePut", func(c *gin.Context) {
        c.String(200, "PUT")
    })

    // Handle DELETE request at /someDelete
    router.DELETE("/someDelete", func(c *gin.Context) {
        c.String(200, "DELETE")
    })

    // Handle PATCH request at /somePatch
    router.PATCH("/somePatch", func(c *gin.Context) {
        c.String(200, "PATCH")
    })

    router.Run(":3000")
}
// 建议自己访问试试吧
/*
http://localhost:3000/someGet
http://localhost:3000/somePost
http://localhost:3000/somePut
http://localhost:3000/someDelete
*/
```

1. **GET**：GET 方法用于获取资源。它只应用于安全和幂等的操作，这意味着多次执行同一 GET 请求应该返回相同的结果，并且不会改变服务器上的资源状态。
2. **POST**：POST 方法用于发送数据到服务器，通常用于创建新的资源。POST 不是幂等的，这意味着多次执行同一 POST 请求可能会得到不同的结果。
3. **PUT**：PUT 方法用于更新资源。它是幂等的，这意味着无论你执行多少次 PUT 请求，结果都是相同的。通常，PUT 请求会包含一个资源的完整更新。
4. **DELETE**：DELETE 方法用于删除资源。它是幂等的，这意味着无论你执行多少次 DELETE 请求，结果都是相同的。

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // 创建一个默认的路由引擎（创建中间件）

	// 创建一个GET路由
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.Query("user")
		pwd := c.Query("pwd") // 这里区分下Param和Query
		c.JSON(200, gin.H{
			"message": "Hello, World! This is a GET request.",
			"id":      id,
			"user":    user,
			"pwd":     pwd,
		})
	})

	// 创建一个POST路由
	r.POST("/path", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! This is a POST request.",
		})
	})

	// 创建一个PUT路由
	r.PUT("/path", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! This is a PUT request.",
		})
	})

	// 创建一个DELETE路由
	r.DELETE("/path/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! This is a DELETE request.",
		})
	})
	// 通过gin生成四种基本HTTP服务

	r.Run(":28081") // 运行服务不加:28081则默认在8080
}
```

显示结果如下：

GET示例：

![image-20240713175920739](./images/image-20240713175920739.png)

DELETE示例；

![image-20240713180550699](./images/image-20240713180550699.png)

其他的自己动手去试试看吧	

## Bind绑定

流程图如下：

![image-20240714103718446](./images/image-20240714103718446.png)

### POST -> Body

```go
package main

import (
	"github.com/gin-gonic/gin"
)

type PostParams struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

// bind绑定的是body体中的数据，如果是get请求，那么就是url中的数据

func main() {
	r := gin.Default()
	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(), // 报错打印err值
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "success",
				"name": p.Name,
				"age":  p.Age,
				"sex":  p.Sex,
			})
		}
	})

	r.Run(":20082")
}

```

结果如下：
![image-20240713182814965](./images/image-20240713182814965.png)

### PUT -> Uri

```go
package main

import (
	"github.com/gin-gonic/gin"
)

type PostParams struct {
	Name string `json:"name" uri:"name"`
	Age  int    `json:"age" uri:"age"`
	Sex  string `json:"sex" uri:"sex"`
}

// 增加（修改）这里同时可以进行PUT绑定

func main() {
	r := gin.Default()
	r.PUT("/testBind/:name/:age/:sex", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindUri(&p)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "success",
				"name": p.Name,
				"age":  p.Age,
				"sex":  p.Sex,
			})
		}
	})

	r.Run(":20082")
}
```

结果如下：

![image-20240713184835618](./images/image-20240713184835618.png)

### PUT -> Url

```go
package main

import (
    "github.com/gin-gonic/gin"
)

type PostParams struct {
    Name string `form:"name" json:"name" uri:"name"`
    Age  int    `form:"age" json:"age" uri:"age"`
    Sex  string `form:"sex" json:"sex" uri:"sex"`
}

func main() {
    r := gin.Default()
    r.PUT("/testBind", func(c *gin.Context) {
        var p PostParams
        err := c.ShouldBindQuery(&p)
        if err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
        } else {
            c.JSON(200, gin.H{
                "msg":  "success",
                "name": p.Name,
                "age":  p.Age,
                "sex":  p.Sex,
            })
        }
    })

    r.Run(":20082")
}
```

结果如下：

![image-20240713184923984](./images/image-20240713184923984.png)

> 可见这里PostParams同时有三张属性即可绑定上面的三种请求
>
> 但是注意路径是/testBind还是/testBind/:name/:age/:sex

### JSON表单验证

```
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"required,mustBig"`
	Sex  string `json:"sex"`
}

func mustBig(f1 validator.FieldLevel) bool {
	if f1.Field().Interface().(int) <= 18 { // reflect类型断言转为int
		return false
	}
	return true
}

//  validator.FieldLevel是一个接口，通常用于自定义验证函数

// 简化如下：
/*
	func mustBig(f1 validator.FieldLevel) bool {
			return f1.Field().Interface().(int) > 18
	}
*/

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mustBig", mustBig)
		}
		// 注册自定义验证函数 先给v注册断言然后绑定标签和方法  主要就是绑定了mustBig标签才能实现精确锁定位置
		var p PostParams
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"data": gin.H{"err": err.Error()}, // gin.H像是将JSON封装返回去
				"msg":  "年龄错误了",
			})
		} else {
			c.JSON(200, gin.H{
				"data": p,
				"msg":  "ok",
			})
		}
	})
	r.Run(":3081")
}
```

正确结果如下：
![image-20240714105742638](./images/image-20240714105742638.png)

验证失败则如下：

![image-20240714105809684](./images/image-20240714105809684.png)

> 简单来说就是年龄值违背了自定义验证器

### 表单验证

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `form:"name"`
	Age  int    `form:"age" binding:"required,mustBig"`
	Sex  string `form:"sex"`
}

func mustBig(f1 validator.FieldLevel) bool {
	return f1.Field().Interface().(int) > 18
}

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mustBig", mustBig)
		}
		var p PostParams
		err := c.ShouldBind(&p)  // 这里是ShouldBind那么就是绑定表单验证了
		if err != nil {
			c.JSON(200, gin.H{
				"data": gin.H{"err": err.Error()},
				"msg":  "年龄错误了",
			})
		} else {
			c.JSON(200, gin.H{
				"data": p,
				"msg":  "ok",
			})
		}
	})
	r.Run(":3081")
}
```

## 文件处理

![image-20240714150816867](./images/image-20240714150816867.png)

![image-20240714150829267](./images/image-20240714150829267.png)

### 用SaveUploadedFile保存文件

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "./"+file.Filename)
		c.JSON(200, gin.H{
			"message": file, // 查看file的信息
		})
	})
	r.Run(":1234")
}
// 因为是自己测试 所以先忽略错误的情况
```

结果如下：
![image-20240714142802797](./images/image-20240714142802797.png)

> 鉴于apipost上传有问题故切换为postman

![image-20240714142956482](./images/image-20240714142956482.png)

> 可见文件被正常保存起来了

### 用IO操作保存文件

修改之后的代码时这样的：
```go
package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)
        // 注意这5行in/out操作
		c.JSON(200, gin.H{
			"message": file, // 查看file的信息
		})
        // 这里先提前说明下这里表示后端返回的时JSON数据
	})
	r.Run(":1234")
}
```

### 后端接收且返回文件

```go
package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)
		c.Writer.Header().Add("Content-Disposition", "attachment; filename="+file.Filename)
		c.File("./" + file.Filename)
        // 这里区别于前面下：返回的是file
	})
	r.Run(":1234")
}
```

结果如下：
![image-20240714145019856](./images/image-20240714145019856.png)

> 解释如下：
>
> 第一行代码 `c.Writer.Header().Add("Content-Disposition", "attachment; filename="+file.Filename)` 是在 HTTP 响应头中添加了一个 "Content-Disposition" 字段。这个字段的值为 "attachment; filename="+file.Filename，意思是告诉客户端，这个响应包含的内容应该被视为一个附件，下载下来并以 file.Filename 为文件名保存。
>
> 第二行代码 `c.File("./" + file.Filename)` 是将文件的内容写入到 HTTP 响应体中。这个文件的路径是当前目录（"./"）加上文件名（file.Filename）。`c.File` 方法会自动设置正确的 `Content-Type` 响应头，根据文件的扩展名来确定文件的 MIME 类型。
>
> > 也就是前面一行是设置响应头 后面一行是设置响应体

![image-20240714145527575](./images/image-20240714145527575.png)

上传图片之类的文件也是可以的

### 多文件处理

```go
```

注意Key依旧是file：

![image-20240714150908227](./images/image-20240714150908227.png)![image-20240714150953547](./images/image-20240714150953547.png)

如果Key不是file那么就会像上面一样出现空括号。

## 路由和分组

