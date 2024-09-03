package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOportunities handles GET requests to retrieve all opportunities
func GetOportunities(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all oportunities",
	})
}

// GetOportunity handles GET requests to retrieve a single opportunity by ID
func GetOportunity(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get opportunity with ID " + id,
	})
}

// CreateOportunity handles POST requests to create a new opportunity
func CreateOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create new opportunity",
	})
}

// UpdateOportunity handles PUT requests to update an existing opportunity by ID
func UpdateOportunity(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update opportunity with ID " + id,
	})
}

// DeleteOportunity handles DELETE requests to delete an opportunity by ID
func DeleteOportunity(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete opportunity with ID " + id,
	})
}
