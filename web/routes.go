package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HelloRoute struct {
}

func NewHelloRoute() *HelloRoute {
	return &HelloRoute{}
}

func (route *HelloRoute) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/hello.html", func(context *gin.Context) {
		fmt.Println("访问：hello.html")
		context.HTML(200, "hello.html", gin.H{
			"title": "nick",
		})
	})

}
