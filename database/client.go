package database

import (
	"OpenSoft-MT/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var dbError error

func ConnectDb(connectionString string){
	dbInstance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if dbError != nil{
		log.Fatal(dbError)
		panic("Error connecting to database!")
	}
	log.Println("Connected to Database!")
}

func Migrate(){
	dbInstance.AutoMigrate(&models.User{})
	log.Println("Database migration complete!")
}

