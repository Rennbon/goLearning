package routers

import (
	"dreamon/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() http.Handler {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/user")
	{
		v1.POST("/login", controllers.LoginHandler)
		v1.POST("/register", controllers.RegisterHandler)
		v1.POST("/unsubscribe", controllers.RemoveUserHandler)
	}

	v2 := router.Group("/wish")
	{
		v2.POST("/addWish", controllers.AddWishHandler)
		v2.POST("/getWish", controllers.GetWishHander)
	}
	fmt.Println("[Plugin Router Profile]...")
	return router
}
