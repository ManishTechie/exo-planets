package engine

import (
	"github.com/gin-gonic/gin"

	"exo-planets/api/v1/controllers"
	"exo-planets/dataservices"
)

// BuildGinEngine creates the Gin Engine with all the middlewares, groups and routes.
func BuildGinEngine(db dataservices.BackendServiceDBInterface, version string) *gin.Engine {
	// create the default Gin engin (GIN_MODE needs to be set beforehand)
	router := gin.New()
	// attach these middlewares at root level, they will apply to every request
	router.Use(
		// recover from panics
		gin.Recovery(),
	)
	// create the /api/v1 sub-router
	v1 := router.Group("/api/v1")
	{
		v1.GET("/healthcheck", controllers.HealthCheck(db))

		v1.POST("/exoplanet", controllers.CreateExoplanet)
		v1.GET("/exoplanet", controllers.GetAllExoplanet)
		v1.GET("/exoplanet/:id", controllers.GetExoplanet)
		v1.PUT("/exoplanet/:id", controllers.UpdateExoplanet)
		v1.DELETE("/exoplanet/:id", controllers.DeleteExoplanet)
		v1.GET("/exoplanet/:id/fuel-estimation/:crew", controllers.CalculateFuelEstimation)

	}
	return router
}
