package database

import(
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"log"
	"fmt"
	"github.com/Shradzz-2111/Task-Manager/models"
)

type Config struct{
	Host string
	Port string
	Password string
	User string
	DBName string
	SSLMode string
}

var DB *gorm.DB

func NewConnection(config *Config) { //(db *gorm.DB, err error) add this if you want to return
	dsn := fmt.Sprintf(
		"host= %s port= %s user= %s password= %s dbname= %s sslmode= %s",
		config.Host,config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil {
		log.Fatal("Not Connected!!")
		return 
	}
	if err := DB.AutoMigrate(&models.Task{}, &models.User{}); err != nil {
          fmt.Errorf("auto-migration failed: %w", err)
		return
    }
	return 
}