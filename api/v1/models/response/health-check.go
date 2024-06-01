package response

type (
	// HealthCheckResponse indicates if the service is healthy, or it failed to check any of the underlying resources.
	HealthCheckResponse struct {
		// Healthy is true if all the checks passed.
		Healthy bool `json:"healthy" example:"false"`

		// Error only included if the services is not healthy.
		Error string `json:"error" example:"Failed to ping the database"`
	}
)
