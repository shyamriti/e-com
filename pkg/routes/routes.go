package routes

import (
	"e-com/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.POST("/user/signup", controller.SignUp)
	r.POST("/user/login", controller.LogIn)
	r.POST("/user/logout", controller.LogOut)
	r.POST("/item", controller.AddItem)
	r.GET("/item", controller.GetItems)
	r.GET("/item/:name", controller.GetItem)
	r.DELETE("/item/:name", controller.DeleteItem)
	return r

}
