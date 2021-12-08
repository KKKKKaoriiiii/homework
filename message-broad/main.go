package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-broad/controller"
	"message-broad/tool"
)

func main() {
	tool.OpenDb()

	r := gin.Default()

	registerRouter(r)

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}

//注册路由
func registerRouter(r *gin.Engine) {
	controller.UserRouter(r)
	controller.MessageRouter(r)
	controller.AdminRouter(r)
}
