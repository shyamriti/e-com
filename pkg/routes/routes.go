package routes

import (
	"e-com/pkg/controller"
	"e-com/pkg/middleware"

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

	users := user.Use(middleware.IsAuthorized)

	users.POST("/logout", controller.LogOut)
	users.GET("/getitems", controller.GetItems)
	users.GET("/search/:name", controller.SearchItem)
	users.GET("/info", controller.UserInfo)
	users.GET("/get", controller.GetUser)
	users.POST("/addcart", controller.AddCart)

	return r

}
