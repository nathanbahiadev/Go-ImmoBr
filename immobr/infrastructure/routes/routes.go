package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanbahiadev/go-immobr/infrastructure/web/controllers"
)

func StartServer() {
	router := gin.Default()

	router.GET("/", controllers.TestController)
	router.POST("/user/", controllers.CreateUserController)
	router.POST("/user/login/", controllers.LoginUserController)

	router.Run()
}
