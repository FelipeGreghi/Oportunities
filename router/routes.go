package router

import (
	"github.com/FelipeGreghi/Oportunities/router/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/oportunities", handler.GetOportunities)
		v1.GET("/oportunities/:id", handler.GetOportunity)
		v1.POST("/oportunities", handler.CreateOportunity)
		v1.PUT("/oportunities/:id", handler.UpdateOportunity)
		v1.DELETE("/oportunities/:id", handler.DeleteOportunity)
	}

}
