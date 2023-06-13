package models

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Item struct {
	ProductName string  `json:"productname"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}

type User struct {
	UserId   int    `json:"userid" gorm:"primary key"`
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	PhoneNo  string `json:"phoneno"`
	Password string `json:"password"`
}

type Order struct {
	User          []User `json:"user" gorm:"foreign key:UserId"`
	OredeId       int    `json:"orderid"`
	ProductDetail []Item `json:"productdetail"`
	Quqantity     int    `json:"quantity"`
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
