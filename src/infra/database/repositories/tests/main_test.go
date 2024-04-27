package repository_test

import (
	"github.com/yuki-katayama/gorm-gin-todo/src/infra/database"
	"github.com/yuki-katayama/gorm-gin-todo/src/domain/models"
    "os"
    "testing"
    "gorm.io/gorm"
    databaseRepo "github.com/yuki-katayama/gorm-gin-todo/src/infra/database/repositories"
    domainRepo "github.com/yuki-katayama/gorm-gin-todo/src/domain/repositories"
)

var db *gorm.DB
var repo domainRepo.TodoRepository

func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}

func setup() {
    var err error
    db, err = database.ConnectionDB();
    if err != nil {
        panic("failed to connect to the database: " + err.Error())
    }
    // データベーススキーマのマイグレーション
    db.AutoMigrate(&models.Todo{})

	// repositoryの初期化
	repo = databaseRepo.NewTodoRepository(db);
    // データベースへのテストデータの挿入（任意）
    seedDatabase()
}

func seedDatabase() {
    todos := []models.Todo{
        {Model: &gorm.Model{}, Content: "Content 1"},
        {Model: &gorm.Model{}, Content: "Content 2"},
    }
    db.Create(&todos)
}

func teardown() {
    // テストが終わった後、テーブルを削除
    db.Migrator().DropTable(&models.Todo{})
}
