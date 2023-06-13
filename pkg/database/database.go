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
	Db, err = gorm.Open(sqlite.Open("new.db"), &gorm.Config{})
	if err != nil {
		panic("failed to database conncetion")
	}
	fmt.Println("Database connected successfully")
	Db.Debug().AutoMigrate(&models.Item{}, &models.User{}, &models.Cart{})

	fmt.Println("Database migrated")
}
