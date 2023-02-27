package controller

import (
	"bubble/model"
	"bubble/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	url --> controller --> service --> model
	请求 --> 控制器 --> 业务逻辑 --> 模型层的增删改查
*/

func IndexHandler(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CretaeOneTodoHandler(ctx *gin.Context) {

	// 1.从请求中拿去数据
	var todo model.Todo
	ctx.ShouldBind(&todo)
	// 2.存入数据库 返回响应
	if err := service.CretaeOneTodo(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func GetTodoAllHandler(ctx *gin.Context) {

	if todoList, err := service.GetTodoAll(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func GetTodoOneHandler(ctx *gin.Context) {
}

func ModifyTodoHandler(ctx *gin.Context) {

	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	todo, err := service.GetTodoOne(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.ShouldBind(todo)
	if err = service.ModifyTodo(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteTodoHandler(ctx *gin.Context) {

	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	if err := service.DeleteTodo(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
