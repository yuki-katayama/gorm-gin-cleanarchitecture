package models

import (
	"gorm.io/gorm"
	"errors"
)

type Todo struct {
	*gorm.Model
	Content string `json:"content"`
}

// Validate - Todoのバリデーションルール
func (t *Todo) Validate() error {
    if t.Content == "" {
        return errors.New("content cannot be empty")
    }
    return nil
}
