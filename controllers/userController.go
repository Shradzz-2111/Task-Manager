package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Shradzz-2111/Task-Manager/database"
	"github.com/Shradzz-2111/Task-Manager/models"
	"github.com/Shradzz-2111/Task-Manager/helpers"
	"log"
	"net/http"
	"time"
	// "strconv"
)


func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		var users []models.User
		if err := database.DB.Find(&users).Error; err != nil{
			log.Print("err getting all users")
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}

		c.JSON(200,users)
	}}


func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		var user models.User
		var userId = c.Param("user_id")
		if err := database.DB.Where("id = ?",userId ).Scan(&user).Error; err != nil {
			log.Print("Sus")
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}

		c.JSON(200,user)
	}
}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context){
		var signUpDetails models.SignUpInput
		//Convert data from JSON to signup Struct
		if err := c.ShouldBindJSON(&signUpDetails) ; err != nil {
			log.Print("Not available")
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		//hash passsword
		hashedPassword, err := helpers.HashPassword(signUpDetails.Password)
		if err != nil {
			log.Panic("Error while hashing")
			return
		}
		//checkl if email is already present
		var count int64
		if err := database.DB.Table("users").Where("email = ?", signUpDetails.Email).Count(&count).Error; err != nil {
			log.Print("Error searching for table")
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		if count > 0 {
			log.Print("User already exists!")
			c.JSON(http.StatusBadRequest,gin.H{"error":"User already exist with given email","message": "User already exists"})
			return
		}
		//create details for User Struct
		now := time.Now()
		newUser :=  models.User{
			FirstName: signUpDetails.FirstName,
			LastName: signUpDetails.LastName,
			Email: signUpDetails.Email,
			Password: hashedPassword,
			CreatedAt: now,
		}

		//inset into database

		if err = database.DB.Create(&newUser).Error; err != nil {
			log.Panic("User not Created")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		} 

		

		// return and send details back 

		c.JSON(http.StatusCreated,gin.H{
			"message":"User Created Sucessfully",
		})
	}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var logInDetails models.LogInInput
		
		if err := c.ShouldBindJSON(&logInDetails); err != nil {
			log.Print("can't login")
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		var user models.User
		if err := database.DB.Find(&user, "email = ?", logInDetails.Email).Error; err != nil {
			log.Print("could't find user")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		if err := helpers.VerifyPassword(user.Password, logInDetails.Password); err != nil{
			log.Print("Wrong Credentials")
			c.JSON(http.StatusBadRequest, gin.H{"error":"Wrong password !"})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"message":"logged in Page",
		})
	}
}

