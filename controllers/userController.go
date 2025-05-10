package controllers

import (
	"github.com/gin-gonic/gin"
)

// config := &database.Config{
// 			Host: os.Getenv("DB_HOST"),
// 			Port: os.Getenv("DB_PORT"),
// 			Password: os.Getenv("DB_PASS"),
// 			User: os.Getenv("DB_USER"),
// 			SSLMode: os.Getenv("DB_SSLMODE"),
// 			DBName: os.Getenv("DB_NAME"),
// 		}

// db, err := database.NewConnection(config)

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		

		c.JSON(200,gin.H{
			"message":"getting all Users",
		})
	}}


func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"getting User",
		})
	}
}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"Signup page",
		})
	}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"login Page",
		})
	}
}

func HashPassword(password string) string{
	return "Hass Pass"
}

func VerifyPassword(userPassword string,providedPassword string) (bool, string){
	return true, "verifying"
}