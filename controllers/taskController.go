package controllers

import (
	// "time"
	// "context"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Shradzz-2111/Task-Manager/models"
	// "gorm.io/gorm"
	"github.com/Shradzz-2111/Task-Manager/database"
)

// type TaskControl struct {
// 	DB *gorm.DB
// }

// func NewTaskController( DB *gorm.DB) TaskControl{
// 	return TaskControl{DB}
// }




func GetTasks() gin.HandlerFunc{
	return func(c *gin.Context){
		// var ctx, cancel = context.WithTimeout(context.Background(),10*time.Second)
		// defer cancel
		config := &database.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}
		db, err := database.NewConnection(config)
		if err != nil {
			log.Print("not connected %s", err)
		}
		var allTasks []models.Task
		if err := db.Raw("SELECT * FROM tasks").Scan(&allTasks).Error; err != nil{
			log.Print("no data")
			c.JSON(402,gin.H{"error":err.Error()})
			return 
		}
		// c.JSON(200,gin.H{
		// 	"message":"getting all Tasks",
		// })
		c.JSON(200,allTasks)
	}
}

func GetTask() gin.HandlerFunc{
	return func(c *gin.Context){
		// var ctx, cancel = context.WithTimeout(context.Background(),10*time.Second)
		// task_id = c.Param("id")
		// var task models.Task
		c.JSON(200,gin.H{
			"message":"getting all Tasks",
		})
	}
}

func CreateTask() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"Creating Tasks",
		})
	}
}

func UpdateTask() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"Updating Tasks",
		})
	}
}

func DeleteTask() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"Deleting Task",
		})
	}
}