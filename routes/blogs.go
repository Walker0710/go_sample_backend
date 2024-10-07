package routes

import (
	"github.com/gin-gonic/gin"
	"Backend/controllers"
	"Backend/middlewares"
)

func BlogRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetBlogs)
	router.GET("/:id", controllers.GetBlog)
	router.POST("/", middlewares.Auth(), controllers.CreateBlog)
	router.POST("/:id/comments", middlewares.Auth(), controllers.CommentOnBlog)
}
