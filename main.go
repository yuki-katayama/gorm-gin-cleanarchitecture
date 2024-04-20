package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-katayama/gorm-gin-todo/models"

	// "github.com/yuki-katayama/gorm-gin-todo/testdata"

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


func main() {
	engine := gin.Default()
	db, err := connectionDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	engine.Static("/static", "./static");
	engine.LoadHTMLGlob("client/*")
	engine.GET("/index", func(c *gin.Context) {
		var todos []models.Todo

		// Get all records
		result := db.Find(&todos)
		if result.Error != nil {
			log.Printf("Error fetching todos: %v", result.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello world",
			"todos": todos,
		})
	})

	//todo edit
	engine.GET("/todos/edit", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatalln(err)
		}
		var todo models.Todo
		db.Where("id = ?", id).Take(&todo)
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"content": "Todo",
			"todo":  todo,
		})
	})

	engine.GET("/todos/destroy", func(c *gin.Context) {
		id, _ := c.GetQuery("id")
		db.Delete(&models.Todo{}, id)
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	engine.POST("/todos/update", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		content := c.PostForm("content")
		var todo models.Todo
		db.Where("id = ?", id).Take(&todo)
		todo.Content = content
		db.Save(&todo)
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	engine.POST("/todos/create", func(c *gin.Context) {
		content := c.PostForm("content")
		db.Create(&models.Todo{Content: content})
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	fmt.Println("Database connection and setup successful")
	engine.Run(":8080")
}
