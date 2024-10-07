package routes

import (
	"github.com/gin-gonic/gin"
	"Backend/controllers"
	"Backend/middlewares"
)

func OverflowRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetOverflows)
	router.GET("/:id", controllers.GetOverflow)
	router.GET("/user/:username", controllers.GetUserOverflows)
	router.POST("/", middlewares.Auth(), controllers.CreateOverflow)
	router.POST("/:id/comments", middlewares.Auth(), controllers.CommentOnOverflow)
}
