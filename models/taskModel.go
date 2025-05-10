package models

import(
	"time"
	"gorm.io/gorm"
)

type Status string

const(
	PENDING   	Status =  "Pending"
	INPROGESS 	Status =  "In-Progress"
	COMPLETED 	Status =  "Completed"
	OVERDUE 	Status =  "Overdue"
)

type Task struct {
	Id 				uint 		`gorm:"primary key;autoIncrement" json:"id"`
	User_ID 		uint 		`json:"user_id"`
	Title 			*string 	`json:"title"`  //* means we need to provide
	Description 	*string 	`json:"description"`
	Status 			Status 		`json:"status"` 
	DueDate 		*time.Time 	`json:"due_date"`
}

func MigrateTasks( db *gorm.DB) error{
	err := db.AutoMigrate(&Task{})
	return err
}