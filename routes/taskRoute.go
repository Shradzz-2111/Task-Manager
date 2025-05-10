package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/Shradzz-2111/Task-Manager/controllers"
)

// type TaskRouteController struct {
// 	TaskController controller.TaskController
// }

// func NewRouteTaskController(taskController controller.TaskController) TaskRouteController{
// 	return TaskRouteController{TaskController}
// }

func TaskRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tasks",controller.GetTasks())
	incomingRoutes.GET("/tasks/:task_id",controller.GetTask())
	incomingRoutes.POST("/tasks",controller.CreateTask())
	incomingRoutes.PATCH("/tasks/:task_id",controller.UpdateTask())
	incomingRoutes.DELETE("/tasks/:task_id",controller.DeleteTask())
}