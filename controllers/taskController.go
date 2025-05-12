package controllers

import (
	"strconv"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Shradzz-2111/Task-Manager/models"
	"github.com/Shradzz-2111/Task-Manager/database"
)



func GetTasks() gin.HandlerFunc{
	return func(c *gin.Context){
		var tasks []models.Task
        currentUser := c.MustGet("currentUser").(models.User) 

		if err := database.DB.Where("user_id =?",currentUser.ID).Find(&tasks).Error; err != nil{
			log.Print("no data")
			c.JSON(402,gin.H{"error":err.Error()})
			return 
		}
		
		
		c.JSON(200,tasks)
	}
}

func GetTask() gin.HandlerFunc{
	return func(c *gin.Context){
    currentUser := c.MustGet("currentUser").(models.User)

	taskId := c.Param("task_id")
		var taskById models.Task

		if err := database.DB.Table("tasks").Where("id = ? and user_id = ?",taskId, currentUser.ID).Scan(&taskById).Error; err != nil{
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
        currentUser := c.MustGet("currentUser").(models.User)

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}
		task.Id = 0
		task.User_ID = currentUser.ID
		if err := database.DB.Create(&task).Error; err !=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}


		c.JSON(http.StatusOK,task)
	}
}

func UpdateTask() gin.HandlerFunc{
	return func(c *gin.Context){
		currentUser := c.MustGet("currentUser").(models.User)
		var taskId = c.Param("task_id")
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}
		if err := database.DB.Model(&models.Task{}).Where("id = ? AND user_id = ?",taskId, currentUser.ID).Updates(&task).Error; err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		var updatedTask models.Task
    	database.DB.Where("id = ?", taskId).First(&updatedTask)
    	c.JSON(http.StatusOK, updatedTask)
	}
}

func DeleteTask() gin.HandlerFunc {
  return func(c *gin.Context) {
    currentUser := c.MustGet("currentUser").(models.User)
    taskId := c.Param("task_id")
	log.Print(taskId)
    taskIDInt, err := strconv.Atoi(taskId)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
      return
    }

    result := database.DB.Where("id = ? AND user_id = ?", taskIDInt, currentUser.ID).Delete(&models.Task{})
    if result.Error != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
  }
}
