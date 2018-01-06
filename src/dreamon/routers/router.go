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
	}
	fmt.Println("[Plugin Router Profile]...")
	return router
}
