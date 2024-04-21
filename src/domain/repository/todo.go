package repository

import (
	"context"
	"github.com/yuki-katayama/gorm-gin-todo/src/domain/models"
)

// TodoRepository - Todoリポジトリのインターフェース
type TodoRepository interface {
	GetByID(ctx context.Context, id uint) (*models.Todo, error)
	Create(ctx context.Context, todo *models.Todo) error
	Update(ctx context.Context, todo *models.Todo) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*models.Todo, error)
}
