package dataservices

import (
	"exo-planets/api/v1"
	"exo-planets/dataservices/model"

	"gorm.io/gorm"
)

// WPDBClient is a wrapper around the default sql.DB object,
// providing helper methods for connecting  and querying a MS-SQL database.
type DBClient struct {
	DB *gorm.DB
}

type BackendServiceDBInterface interface {
	Connect(connectionString string) (setupError error)
	Ping() error
	Close() error

	BeginTransaction() (txn BackendServiceDBInterface, err *api.APIError)
	CommitTransaction() (err *api.APIError)
	RollbackTransaction() (err *api.APIError)

	CreateExoplanet(exoplanet model.Exoplanet) (err error)
	GetAllExoplanet() (exoplanets []model.Exoplanet, err error)
	GetExoplanetByID(id string) (exoplanet *model.Exoplanet, err error)
	UpdateExoplanetByID(id string, exoplanet model.Exoplanet) (err error)
	DeleteExoplanetByID(id string) (err error)
}
