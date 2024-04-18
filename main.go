package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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



func connectionDB() (*gorm.DB, error) {
	config := getDBConfig();
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Table)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

type Todo struct {
	*gorm.Model
	Name string
	Title string
}

func main() {
	// r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r := gin.Default()
	db, err := connectionDB();
	if err != nil {
		log.Fatal(err);
		return
	}
	// dbをmigrateします
	db.AutoMigrate(&Todo{})
	fmt.Println(db)
	r.Run(":8080")
}
