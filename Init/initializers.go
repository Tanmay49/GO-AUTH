package Init

import (
	"fmt"
	"log"
	"main/Model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=tiny.db.elephantsql.com user=iwdokbkl password=vYpJpjtx8XMrkWikZ26KfJVD7lBNXEQO dbname=iwdokbkl port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&Model.User{})
}
