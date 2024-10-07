package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors" // Import the CORS package
	"Backend/config"
	"Backend/routes"
)

func main() {
	// Load the configuration and connect to the database
	config.LoadConfig()
	config.ConnectDB()

	// Create a new Gin router
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Replace with your front-end URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"}, // Expose these headers
		AllowCredentials: true, // Allow credentials (cookies, authorization headers, etc.)
		MaxAge: 12 * 3600, // Maximum age of preflight request cache
	}))

	// Define API routes
	api := router.Group("/api")
	routes.AuthRoutes(api.Group("/auth"))
	routes.BlogRoutes(api.Group("/blogs"))
	routes.OverflowRoutes(api.Group("/overflows"))
	routes.UserRoutes(api.Group("/users"))

	// Start the server
	router.Run(":5000")
}
