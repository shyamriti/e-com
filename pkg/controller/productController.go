package controller

import (
	"e-com/pkg/database"
	"e-com/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddItem(c *gin.Context) {
	var item models.Item
	if err := c.BindJSON(&item); err != nil {
		log.Println(err)
	}
	database.Db.Create(&item)
	c.JSON(200, item)
}

func SearchItem(c *gin.Context) {
	var item models.Item

	err := database.Db.Where("name= ?", c.Param("name")).First(&item)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, item)
}

func GetItems(c *gin.Context) {
	var item []models.Item
	resp := database.Db.Find(&item)
	if resp.Error != nil {
		log.Println(resp.Error)
	}
	c.JSON(200, item)
}
func DeleteItem(c *gin.Context) {
	var item models.Item

	name := c.Param("name")
	database.Db.Where("name= ?", name).Delete(&item)
	c.JSON(200, "data deleted")
}

func AddCart(c *gin.Context) {
	var item models.Item
	itemIDString := c.Query("itemid")
	itemID, err := strconv.Atoi(itemIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Id has to be an integer")
		return
	}
	err = database.Db.Where("id", itemID).First(&item).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Id not found")
		return
	}
	err = database.Db.Debug().Create(&models.Cart{
		ItemID: uint(item.ID),
	}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Unable to insert item to cart")
		return
	}
	c.JSON(200, "item added to cart")
}
