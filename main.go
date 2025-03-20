package main

import (
	"fmt"
	"gin_demo2/routers"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("2-我是一个中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func main() {
	r := gin.Default()

	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})

	// 加载所有模板文件
	r.LoadHTMLGlob("templates/*.html")

	// 配置静态文件路径
	r.Static("/static", "./static")

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test")
	})
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	//Get请求传值
	r.GET("/article", func(c *gin.Context) {

		id := c.DefaultQuery("id", "1")

		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})
	//post演示
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{})
	})

	//获取表单post过来的数据
	r.POST("/doAddUser1", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	r.POST("/doAddUser2", func(c *gin.Context) {
		user := UserInfo{}
		err := c.ShouldBind(&user)
		if err == nil {
			fmt.Println("%#v", user.Password)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	//动态路由传值
	routers.ApiRouterInit(r)

	r.Run()
}
