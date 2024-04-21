package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/yuki-katayama/gorm-gin-todo/src/infra/database/repository"
	"github.com/yuki-katayama/gorm-gin-todo/src/infra/database"
	"github.com/yuki-katayama/gorm-gin-todo/src/application/usecases"
	"github.com/yuki-katayama/gorm-gin-todo/src/interface/controller"
	"github.com/yuki-katayama/gorm-gin-todo/src/infra/http/routes"
)

func main() {
	// データベース接続の設定
	db, err := database.ConnectionDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// リポジトリの初期化
	todoRepo := repository.NewTodoRepository(db)

	// サービス層の初期化
	todoService := service.NewTodoService(todoRepo)

	// コントローラの初期化
	todoController := controllers.NewTodoController(todoService)

	engine := gin.Default()
	// ルータの設定
	engine = router.SetupRouterTodo(engine, todoController)
	engine = router.SetupRouterPage(engine, todoController)

	// サーバを8080ポートで起動
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
