package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"strings"
	"github.com/Shradzz-2111/Task-Manager/helpers"
	"github.com/Shradzz-2111/Task-Manager/models"
	"github.com/Shradzz-2111/Task-Manager/database"
)

func Authentication() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
		}
		claims, err := helpers.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Token"})
			log.Printf(tokenString)
            c.Abort()
            return
		}

		var user models.User
		if err := database.DB.First(&user, claims.UserID).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            c.Abort()
            return
        }
		c.Set("currentUser", user)
		c.Next()
	}
}