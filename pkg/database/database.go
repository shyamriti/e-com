package database

import (
	"e-com/pkg/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connection() {
	var err error
	Db, err = gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to database conncetion")
	}
	fmt.Println("Database connected successfully")
	Db.AutoMigrate(&models.Item{}, &models.Order{}, &models.User{})
	fmt.Println("Database migrated")
}
