package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/Shradzz-2111/Task-Manager/controllers"
	"github.com/Shradzz-2111/Task-Manager/middleware"
)


func TaskRoutes(incomingRoutes *gin.Engine){
	authGroup := incomingRoutes.Group("/")
    authGroup.Use(middleware.Authentication())

    authGroup.GET("/tasks", controller.GetTasks())
    authGroup.GET("/tasks/:task_id", controller.GetTask())
    authGroup.POST("/tasks", controller.CreateTask())
    authGroup.PATCH("/tasks/:task_id", controller.UpdateTask())
    authGroup.DELETE("/tasks/:task_id", controller.DeleteTask())
}