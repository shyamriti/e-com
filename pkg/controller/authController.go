package controller

import (
	"e-com/pkg/auth"
	"e-com/pkg/database"
	"e-com/pkg/models"
	"e-com/pkg/request"
	"log"
	"net/http"

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
		c.AbortWithStatusJSON(http.StatusBadRequest, "binding failed")
		return
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
		SecretKey:        "secret",
		Issuer:           "AuthService",
		ExpirationMinute: 1,
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
	// token, err := jwtWrapper.ValidateToken(signedToken)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, err)
	// }
	c.SetCookie("user", signedToken, 3600, "/", "localhost", false, true)

	cookie, err := c.Cookie("user")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "cookie not set")
		return
	}
	c.JSON(200, cookie)

}

func LogOut(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "localhost", false, true)

	cookie, err := c.Cookie("user")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "log out abort")
		return
	}
	if cookie != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "cookies not nil")
		return

	}
	c.JSON(200, "log out successfully")
}
