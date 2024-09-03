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

// GetOportunities handles GET requests to retrieve all opportunities or a specific one by ID
func GetOportunities(ctx *gin.Context) {
	id := ctx.Query("id")
	openings := []schemas.Opening{}

	if id != "" {
		if err := db.Where("id = ?", id).Find(&openings).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		if err := db.Find(&openings).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"openings": openings,
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
	id := ctx.Query("id")
	logger.Infof("ID: %v", id)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query Param ID is required",
		})
		return
	}
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating opportunity: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	opening := schemas.Opening{}
	// Find opportunity by ID
	if err := db.First(&opening, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Opportunity not found",
		})
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opportunity updated successfully",
		"opening": opening,
	})
}

// DeleteOportunity handles DELETE requests to delete an opportunity by ID
func DeleteOportunity(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query Param ID is required",
		})
		return
	}

	opening := schemas.Opening{}
	// Find opportunity by ID
	if err := db.First(&opening, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Opportunity not found",
		})
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opportunity deleted successfully",
		"opening": opening,
	})
}
