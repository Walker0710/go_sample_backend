package routes

import (
	"github.com/gin-gonic/gin"
	"Backend/controllers"
	"Backend/middlewares"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/profile", middlewares.Auth(), controllers.Profile)
}


