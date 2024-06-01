package controllers

import (
	"net/http"

	"gas-giant/api/v1/models/response"
	"github.com/gin-gonic/gin"
)

// Pinger is every object with a Ping method that returns an optional error,
// in case the pinger fails to ping the underlying service.
type Pinger interface {
	Ping() error
}

// HealthCheck accepts zero or more pingers and returns a handler, which iterate over the pingers
// and returns 500 InternalServerError if any of them fails, and 200 OK if non of them failed.
// @Summary check the service's health status
// @Description HealthCheck will report healthy status if all the underlying dependencies are available (e.g.: db connection)
// @Tags Health-Check
// @Accept  json
// @Produce  json
// @Success 200 {object} response.HealthCheckResponse
// @Failure 500 {object} response.HealthCheckResponse
// @Router /api/v1/healthcheck [get]
func HealthCheck(pingers ...Pinger) func(c *gin.Context) {
	return func(c *gin.Context) {
		for _, pinger := range pingers {
			if err := pinger.Ping(); err != nil {

				c.JSON(http.StatusInternalServerError, response.HealthCheckResponse{
					Healthy: false,
					Error:   err.Error(),
				})
				return
			}
		}
		c.JSON(http.StatusOK, response.HealthCheckResponse{Healthy: true})
	}
}
