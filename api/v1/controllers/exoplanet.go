package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"exo-planets/api/v1/models/request"
	"exo-planets/api/v1/models/response"
	"exo-planets/dataservices"
	"exo-planets/logging"
	"exo-planets/util"
)

// POST /api/v1/exoplanet
// CreateExoplanet -
func CreateExoplanet(c *gin.Context) {
	var exoplanet request.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := exoplanet.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exoplanet.ID = util.GenerateUUID()
	if err := dataservices.DB().DB.Create(&exoplanet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exoplanet)
}

// GET /api/v1/exoplanet
// GetAllExoplanet -
func GetAllExoplanet(c *gin.Context) {
	var exoplanets []request.Exoplanet

	result := dataservices.DB().DB.Find(&exoplanets)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if len(exoplanets) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "exoplanet not found"})
		return
	}

	c.JSON(http.StatusCreated, exoplanets)

}

// GET /api/v1/exoplanet/{id}
// GetAllExoplanet -
func GetExoplanet(c *gin.Context) {
	logger := logging.GetLogger()

	idStr := c.Param("id")
	var exoplanet response.Exoplanet

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("Invalid UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	result := dataservices.DB().DB.Where("id=?", id.String()).First(&exoplanet)
	if result.Error != nil {
		logger.Error("error while fetching details from db", zap.Error(result.Error), zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if exoplanet.ID == nil {
		logger.Error("not found", zap.String("id", id.String()))
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusCreated, exoplanet)

}
