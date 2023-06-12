package controller

import (
	"e-com/pkg/auth"
	"e-com/pkg/database"
	"e-com/pkg/models"
	"e-com/pkg/request"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
