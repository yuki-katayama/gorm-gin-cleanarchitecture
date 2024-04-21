package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/yuki-katayama/gorm-gin-todo/src/interface/controllers"
)

// SetupRouter - ルータの設定を行います
func SetupRouterTodo(engine *gin.Engine, todoController *controllers.TodoController) *gin.Engine {
    // HTMLテンプレートのロード

	engine.GET("/todo/list", func(c *gin.Context) {
		todos, err := todoController.ListTodos(c)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve todos"})
            return
        }
        c.JSON(http.StatusOK, todos)
	})
	engine.GET("/todo/get", func(c *gin.Context) {
		todo, err := todoController.GetTodoByID(c)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})
	engine.POST("/todo/create", todoController.CreateTodo)
	engine.POST("/todo/update", todoController.UpdateTodo)
	engine.GET("/todo/delete", todoController.DeleteTodo)

	return engine
}
