package handler

import (
	"net/http"

	"github.com/FelipeGreghi/Oportunities/config"
	"github.com/FelipeGreghi/Oportunities/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("Handler")
	db = config.GetSQLite()
}

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
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating opportunity: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opportunity: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Opportunity created successfully",
		"opening": opening,
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
