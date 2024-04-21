package repository_test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/yuki-katayama/gorm-gin-todo/src/domain/models"
)

// Todoが正しく取得できるかテスト
func TestGetByID(t *testing.T) {
    ctx := context.Background()

    // ID 1がseedデータとして存在することを前提とする
    todo, err := repo.GetByID(ctx, 1)
    assert.Nil(t, err)
    assert.NotNil(t, todo)
    assert.Equal(t, uint(1), todo.ID)
    assert.Equal(t, "Content 1", todo.Content)
}

// Todoが正しく作成できるかテスト
func TestCreate(t *testing.T) {
    ctx := context.Background()
    todo := &models.Todo{Content: "Valid New Todo"}

    // バリデーションチェック
    err := todo.Validate()
    assert.Nil(t, err)

    // Todoの作成
    if err == nil {
        err = repo.Create(ctx, todo)
        assert.Nil(t, err)
        assert.NotZero(t, todo.ID)
    }
}

// Todoの更新テスト
func TestUpdate(t *testing.T) {
    ctx := context.Background()

    // 既存のTodoを取得
    todo, _ := repo.GetByID(ctx, 1)
    originalContent := todo.Content
    todo.Content = "Updated Content"

    // 更新処理
    err := repo.Update(ctx, todo)
    assert.Nil(t, err)
    assert.NotEqual(t, originalContent, todo.Content)
}

// Todoの削除テスト
func TestDelete(t *testing.T) {
    ctx := context.Background()

    // 新しいTodoを作成し、それを削除
    newTodo := &models.Todo{Content: "To be deleted"}
    repo.Create(ctx, newTodo)

    err := repo.Delete(ctx, newTodo.ID)
    assert.Nil(t, err)
}

// Todoリストが正しく取得できるかテスト
func TestList(t *testing.T) {
    ctx := context.Background()

    // Todoリストを取得
    todos, err := repo.List(ctx)
    assert.Nil(t, err)
    assert.NotEmpty(t, todos)
    assert.True(t, len(todos) >= 2) // seedデータが2つ存在することを前提
}
