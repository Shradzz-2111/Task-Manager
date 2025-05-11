package models

import(
	"time"
	"gorm.io/gorm"
	
)

type User struct{
	ID 			uint `gorm:"primary key;autoIncrement" json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	CreatedAt 	time.Time `json:"created_at"`

}

func MigrateUsers(db *gorm.DB) error{
	err := db.AutoMigrate(&User{})
	return err
}

type SignUpInput struct{
	FirstName 		string `json:"first_name" binding:"required"` 
	LastName 		string `json:"last_name" binding:"required"`
	Email 			string `json:"email" binding:"required"`
	Password 		string `json:"password" binding:"required"`
}

type LogInInput struct {
	Email 			string `json:"email" binding:"required"`
	Password 		string `json:"password" binding:"required"`	
}