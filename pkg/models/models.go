package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
}

type Cart struct {
	gorm.Model
	ItemID uint
	Item   []Item `json:"item" gorm:"foreignKey:ID;association_foreignkey:ItemID"`
}

type User struct {
	gorm.Model
	UserName     string `json:"user_name"`
	Email        string `json:"email" gorm:"unique"`
	PhoneNo      string `json:"phone_no"`
	Password     string `json:"password"`
	IsAmbassador bool   `json:"is_ambassador"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)

	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		log.Println(err)

	}
	return nil
}
