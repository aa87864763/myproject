package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type prince struct {
	Name   string
	Age    int
	Nation string
}

func sayhello(c *gin.Context) {

	wyz := prince{
		Name:   "万永智",
		Age:    25,
		Nation: "CN",
	}
	hobbylist := []string{
		"篮球",
		"麻将",
		"双色球",
	}
	data := map[string]interface{}{
		"prince1": wyz,
		"hobby":   hobbylist,
	}
	c.HTML(http.StatusOK, "default/index.html", data)
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", sayhello)
	r.Run()
}
