package repository

import (
	"context"
	"gorm.io/gorm"
	"github.com/yuki-katayama/gorm-gin-todo/src/domain/models"
	"github.com/yuki-katayama/gorm-gin-todo/src/domain/repository"
)

// TodoRepository - GORMによるTodoリポジトリの実装
type TodoRepository struct {
	DB *gorm.DB
}

// NewTodoRepository - 新しいGormTodoRepositoryを作成します
func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &TodoRepository{DB: db}
}

// 以下、インターフェースの各メソッドの実装
func (r *TodoRepository) GetByID(ctx context.Context, id uint) (*models.Todo, error) {
	var todo models.Todo
	result := r.DB.First(&todo, id)
	return &todo, result.Error
}

func (r *TodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) Update(ctx context.Context, todo *models.Todo) error {
	return r.DB.WithContext(ctx).Save(todo).Error
}

func (r *TodoRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.Delete(&models.Todo{}, id).Error
}

func (r *TodoRepository) List(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	result := r.DB.Find(&todos)
	return todos, result.Error
}
