package controller

import (
	"e-com/pkg/database"
	"e-com/pkg/middleware"
	"e-com/pkg/models"
	"log"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var user models.User
	id, _ := middleware.GetUserID(c)
	database.Db.Where("userid= ", &id).First(&user)
	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	var user []models.User
	resp := database.Db.Find(&user)
	if resp.Error != nil {
		log.Println(resp.Error)
	}
	c.JSON(200, user)
}
