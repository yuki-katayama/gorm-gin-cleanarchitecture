package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yuki-katayama/gorm-gin-todo/src/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	User string
	Password string
	Host string
	Port int
	Table string
}

// OpenDB - データベース接続を開きます
func getDBConfig() DBConfig {
    port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
    return DBConfig{
        User: os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        Host: os.Getenv("DB_HOST"),
        Port: port,
		Table: os.Getenv("DB_TABLE"),
    }
}

func ConnectionDB() (*gorm.DB, error) {
	config := getDBConfig();
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Table)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		return nil, err
	}
	return db, nil
}