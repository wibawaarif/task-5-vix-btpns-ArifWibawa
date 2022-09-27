package database

import (
	"FinalProject/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("${USERNAME}:${PASSWORD@tcp(localhost:${PORT})/${DATABASE_NAME}?parseTime=true"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.User{})

	DB = database
}