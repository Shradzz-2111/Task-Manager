package controllers

import (
	// "time"
	// "context"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Shradzz-2111/Task-Manager/models"
	"github.com/Shradzz-2111/Task-Manager/database"
)



func GetTasks() gin.HandlerFunc{
	return func(c *gin.Context){
		var tasks []models.Task


		if err := database.DB.Find(&tasks).Error; err != nil{
			log.Print("no data")
			c.JSON(402,gin.H{"error":err.Error()})
			return 
		}
		
		c.JSON(200,gin.H{
			"message":"getting all Tasks",
		})
		c.JSON(200,tasks)
	}
}

func GetTask() gin.HandlerFunc{
	return func(c *gin.Context){
	taskId := c.Param("task_id")
		var taskById models.Task

		if err := database.DB.Table("tasks").Where("id = ?",taskId).Scan(&taskById).Error; err != nil{
			log.Print(err)
			c.JSON(404,gin.H{"message":"Not Found"})
			return
		}
		c.JSON(200,taskById)
	}
}

func CreateTask() gin.HandlerFunc{
	return func(c *gin.Context){
		var task models.Task

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}
		task.Id = 0
		if err := database.DB.Create(&task).Error; err !=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}


		c.JSON(http.StatusOK,task)
	}
}

func UpdateTask() gin.HandlerFunc{
	return func(c *gin.Context){
		var taskId = c.Param("task_id")
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}
		if err := database.DB.Model(&models.Task{}).Where("id = ?",taskId).Updates(&task).Error; err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		c.JSON(200,gin.H{
			"message":"Updated Tasks",
		})
	}
}

func DeleteTask() gin.HandlerFunc{
	return func(c *gin.Context){
		var taskId = c.Param("task_id")
		if err := database.DB.Delete(&models.Task{},taskId).Error; err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{"message":"Task deleted sucessfully!"})
	}
}