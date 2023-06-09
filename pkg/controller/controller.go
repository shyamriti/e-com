package controller

import (
	"e-com/pkg/auth"
	"e-com/pkg/database"
	"e-com/pkg/models"
	"e-com/pkg/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

var Item models.Item

func SignUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
	}
	database.Db.Create(&user)
	c.JSON(200, user)

}

func LogIn(c *gin.Context) {
	var payload request.LoginPayload
	var user models.User

	err := c.ShouldBind(&payload)
	if err != nil {
		log.Println(err)
	}
	result := database.Db.Where("email=?", payload.Email).Select("*").First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}
	cookie, err := c.Cookie("user")
	if err != nil {
		c.SetCookie("user", signedToken, 3600, "/", "127.0.0.1", false, true)
		c.JSON(200, gin.H{"msg": "cookie set successfully"})
	} else {
		c.JSON(200, gin.H{"msg": cookie})
	}

}

func LogOut(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.SetCookie("user", "", -1, "/", "127.0.0.1", false, true)
		c.JSON(200, gin.H{"msg": "logout successfully"})
	} else {
		c.JSON(200, gin.H{"msg": cookie})
	}
}

func AddItem(c *gin.Context) {
	if err := c.BindJSON(&Item); err != nil {
		log.Fatal(err)
	}
	database.Db.Create(&Item)
	c.JSON(200, Item)
}

func GetItem(c *gin.Context) {
	err := database.Db.Where("name= ?", c.Param("name")).First(&Item)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, Item)
}

func GetItems(c *gin.Context) {
	var item []models.Item
	resp := database.Db.Find(&item)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	c.JSON(200, item)
}
func DeleteItem(c *gin.Context) {
	name := c.Param("name")
	database.Db.Where("name= ?", name).Delete(&Item)
	c.JSON(200, "data deleted")
}
