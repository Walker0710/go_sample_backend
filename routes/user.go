package routes

import (
	"github.com/gin-gonic/gin"
	"Backend/controllers"
	"Backend/middlewares"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/profile", middlewares.Auth(), controllers.Profile)
}
