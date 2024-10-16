package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"exo-planets/api/v1/models/request"
	"exo-planets/api/v1/models/response"
	"exo-planets/dataservices"
	"exo-planets/dataservices/model"
	"exo-planets/logging"
	"exo-planets/util"
)

// POST /api/v1/exoplanet
// CreateExoplanet -
func CreateExoplanet(c *gin.Context) {
	var payload request.Exoplanet
	logger := logging.GetLogger()

	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error("error while binding the request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := payload.Validate(); err != nil {
		logger.Error("error while valdating the request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exoplanet := model.Exoplanet{
		ID:          util.GenerateUUID(),
		Name:        payload.Name,
		Description: payload.Description,
		Distance:    payload.Distance,
		Radius:      payload.Radius,
		Mass:        payload.Mass,
		Type:        (*model.ExoplanetsType)(payload.Type),
	}

	if err := dataservices.DB().CreateExoplanet(exoplanet); err != nil {
		logger.Error("error while creating the exoplanet", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("exoplanet created")
	c.JSON(http.StatusCreated, payload)
}

// GET /api/v1/exoplanet
// GetAllExoplanet -
func GetAllExoplanet(c *gin.Context) {

	logger := logging.GetLogger()

	exoplanets, err := dataservices.DB().GetAllExoplanet()
	if err != nil {
		logger.Error("error while getting details from DB", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("Invalid UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	exoplanet, err := dataservices.DB().GetExoplanetByID(id.String())
	if err != nil {
		logger.Error("error while fetching details from db", zap.Error(err), zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("exoplanet found")
	c.JSON(http.StatusCreated, exoplanet)

}

// PUT /api/v1/exoplanet/:id
// UpdateExoplanet -
func UpdateExoplanet(c *gin.Context) {
	logger := logging.GetLogger()
	var payload request.Exoplanet

	if err := c.ShouldBindJSON(&payload); err != nil {
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
	_, err = dataservices.DB().GetExoplanetByID(id.String())
	if err != nil {
		logger.Error("error while fetching details from db", zap.Error(err), zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payloadUpdate := model.Exoplanet{
		Name:        payload.Name,
		Description: payload.Description,
		Distance:    payload.Distance,
		Radius:      payload.Radius,
		Mass:        payload.Mass,
		Type:        (*model.ExoplanetsType)(payload.Type),
	}

	if err := dataservices.DB().UpdateExoplanetByID(id.String(), payloadUpdate); err != nil {
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
	_, err = dataservices.DB().GetExoplanetByID(id.String())
	if err != nil {
		logger.Error("error while fetching details from db", zap.Error(err), zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = dataservices.DB().DeleteExoplanetByID(id.String())
	if err != nil {
		logger.Error("error while deleting the exoplanet", zap.String("id", id.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	logger.Info("exoplanet deleted successfully")

	c.JSON(http.StatusOK, map[string]string{
		"message": "exoplanet deleted successfully",
	})
}

// GET /exoplanet/:id/fuel-estimation
// CalculateFuelEstimation -
func CalculateFuelEstimation(c *gin.Context) {

	logger := logging.GetLogger()

	idStr := c.Param("id")
	crewStr := c.Param("crew")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("Invalid UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	crew, err := strconv.ParseFloat(crewStr, 64)
	if err != nil {
		logger.Error("invalid crew no", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Crewn no"})
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

	distance := getExoplanet.Distance
	radius := getExoplanet.Radius

	var gravity float64

	if getExoplanet.Type == request.GasGiantType.String() {
		gravity = 0.5 / math.Pow(radius, 2)
	} else {
		gravity = getExoplanet.Mass / math.Pow(radius, 2)
	}

	fuelEstimation := distance / math.Pow(gravity, 2) * crew

	c.JSON(http.StatusOK, map[string]float64{
		"fuel-estimation": fuelEstimation,
	})
}
