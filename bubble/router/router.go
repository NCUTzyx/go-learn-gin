package router

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "./static")
	// 加载模板文件
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("/v1")
	{
		// 待办事项
		// 添加one
		v1Group.POST("/todo", controller.CretaeOneTodoHandler)

		// 查看all
		v1Group.GET("/todo", controller.GetTodoAllHandler)

		// 查看one
		v1Group.GET("/todo/:id", controller.GetTodoOneHandler)

		// 修改one
		v1Group.PUT("/todo/:id", controller.ModifyTodoHandler)

		// 删除one
		v1Group.DELETE("/todo/:id", controller.DeleteTodoHandler)
	}

	return r

}
