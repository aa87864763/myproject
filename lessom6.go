package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lesson.html", gin.H{
			"name": "小王子",
		})
	})
	r.Run()

}
