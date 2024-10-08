package router

import (
	"github.com/FelipeGreghi/Oportunities/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize the Handler
	handler.Init()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/oportunities", handler.GetOportunities)
		v1.POST("/oportunities", handler.CreateOportunity)
		v1.PUT("/oportunities", handler.UpdateOportunity)
		v1.DELETE("/oportunities", handler.DeleteOportunity)
	}

}
