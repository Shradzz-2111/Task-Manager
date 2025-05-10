package main

import(
	"os"
	"log"
	"github.com/Shradzz-2111/Task-Manager/database"
	"github.com/Shradzz-2111/Task-Manager/routes"
	"github.com/Shradzz-2111/Task-Manager/initializers"
	// "github.com/Shradzz-2111/Task-Manager/models"
	// "github.com/Shradzz-2111/Task-Manager/middleware"
	// "gorm.io/gorm"
	// "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	config := &database.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}

	database.NewConnection(config)
}


func main(){

	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
		log.Printf("Defaulting to port %s",port)
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	routes.TaskRoutes(router)

	router.GET("/ping",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})

	router.Run(":" + port)
}