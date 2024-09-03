package router

import "github.com/gin-gonic/gin"

func Init() {
	// Creates a gin router with default middleware
	router := gin.Default()

	// Initializes the routes
	InitializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
