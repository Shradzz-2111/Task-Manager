package models

import(
	"time"
	"gorm.io/gorm"
)

type User struct{
	ID 			uint `gorm:"primary key;autoIncrement" json:"id"`
	FirstName 	*string `json:"first_name"`
	LastName 	*string `json:"last_name"`
	Email 		*string `json:"email"`
	Pasword 	*string `json:"password"`
	CreatedAt 	time.Time `json:"created_at"`

}

func MigrateUsers(db *gorm.DB) error{
	err := db.AutoMigrate(&User{})
	return err
}