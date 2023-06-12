package routes

import (
	"e-com/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	admin := r.Group("/admin")
	admin.POST("/signup", controller.SignUp)
	admin.POST("/login", controller.LogIn)
	admin.POST("/logout", controller.LogOut)
	admin.POST("/additem", controller.AddItem)
	admin.GET("/getitems", controller.GetItems)
	admin.GET("/search/:name", controller.SearchItem)
	admin.DELETE("/deleteitem/:name", controller.DeleteItem)

	user := r.Group("/user")
	user.POST("/signup", controller.SignUp)
	user.POST("/login", controller.LogIn)
	user.POST("/logout", controller.LogOut)
	user.GET("/getitems", controller.GetItems)
	user.GET("/search/:name", controller.SearchItem)
	user.GET("/info", controller.UserInfo)
	user.GET("/get", controller.GetUser)
	user.POST("/addcart", controller.AddCart)

	return r

}
