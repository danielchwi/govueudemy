package database

import (
	"govue/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	database, err := gorm.Open(mysql.Open("root:@/govue"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to databases")
	}

	database.AutoMigrate(&models.User{})
}
