package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuki-katayama/gorm-gin-todo/src/interface/controllers"
)

// SetupRouter - ルータの設定を行います
func SetupRouterTodo(engine *gin.Engine, todoController *controllers.TodoController) *gin.Engine {
    // HTMLテンプレートのロード

	engine.POST("/todos/create", todoController.CreateTodo)
	engine.POST("/todos/update", todoController.UpdateTodo)
	engine.GET("/todos/destroy", todoController.DeleteTodo)

	return engine
}