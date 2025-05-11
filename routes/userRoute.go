package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/Shradzz-2111/Task-Manager/controllers"
	"github.com/Shradzz-2111/Task-Manager/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
	incomingRoutes.POST("/users/signup",controller.SignUp())
	incomingRoutes.POST("/users/login",controller.Login())
	authGroup := incomingRoutes.Group("/")
    authGroup.Use(middleware.Authentication())
    {
        authGroup.GET("/me", controller.GetMe())
    }
}