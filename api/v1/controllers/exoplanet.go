package controllers

import (
	"net/http"
	"time"

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
	logger := logging.GetLogger()

	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		logger.Error("error while binding the request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := exoplanet.Validate(); err != nil {
		logger.Error("error while valdating the request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exoplanet.ID = util.GenerateUUID()
	if err := dataservices.DB().DB.Create(&exoplanet).Error; err != nil {
		logger.Error("error while creating the exoplanet", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("exoplanet created")
	c.JSON(http.StatusCreated, exoplanet)
}

// GET /api/v1/exoplanet
// GetAllExoplanet -
func GetAllExoplanet(c *gin.Context) {
	var exoplanets []request.Exoplanet

	logger := logging.GetLogger()

	result := dataservices.DB().DB.Where("deleted_at IS NULL").Find(&exoplanets)
	if result.Error != nil {
		logger.Error("error while getting details from DB", zap.Error(result.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if len(exoplanets) == 0 {
		logger.Error("not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "exoplanet not found"})
		return
	}
	logger.Info("exoplanets found", zap.Int("total no of exoplanets", len(exoplanets)))
	c.JSON(http.StatusCreated, exoplanets)

}

// GET /api/v1/exoplanet/:id
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

	result := dataservices.DB().DB.Where("id=? and deleted_at IS NULL", id.String()).First(&exoplanet)
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
	logger.Info("exoplanet found")
	c.JSON(http.StatusCreated, exoplanet)

}

// PUT /api/v1/exoplanet/:id
// UpdateExoplanet -
func UpdateExoplanet(c *gin.Context) {
	logger := logging.GetLogger()
	var exoplanet request.Exoplanet

	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		logger.Error("error while binding the request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("Invalid UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	getExoplanet := response.Exoplanet{}
	result := dataservices.DB().DB.Where("id=?", id.String()).First(&getExoplanet)
	if result.Error != nil {
		logger.Error("error while fetching details from db", zap.Error(result.Error), zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if getExoplanet.ID.String() == "" {
		logger.Error("not found", zap.String("id", id.String()))
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	if err := dataservices.DB().DB.Where("id=?", id.String()).Omit("id").Updates(&exoplanet).Error; err != nil {
		logger.Error("error while updating exoplanet details", zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	logger.Info("exoplanet updated successfully")

	c.JSON(http.StatusOK, map[string]string{
		"message": "exoplanet updated successfully",
	})
}

// PUT /api/v1/exoplanet/:id
// DeleteExoplanet -
func DeleteExoplanet(c *gin.Context) {

	logger := logging.GetLogger()

	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("Invalid UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	getExoplanet := response.Exoplanet{}
	result := dataservices.DB().DB.Where("id=?", id.String()).First(&getExoplanet)
	if result.Error != nil {
		logger.Error("error while fetching details from db", zap.Error(result.Error), zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if getExoplanet.ID.String() == "" {
		logger.Error("not found", zap.String("id", id.String()))
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	result = dataservices.DB().DB.Model(request.Exoplanet{}).Where("id=?", id.String()).Update("deleted_at", time.Now())
	if result.Error != nil {
		logger.Error("error while deleting the exoplanet", zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	logger.Info("exoplanet deleted successfully")

	c.JSON(http.StatusOK, map[string]string{
		"message": "exoplanet deleted successfully",
	})
}
