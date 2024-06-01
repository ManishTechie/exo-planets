package dataservices

import (
	"gas-giant/api/v1"
	"gorm.io/gorm"
)

// WPDBClient is a wrapper around the default sql.DB object,
// providing helper methods for connecting  and querying a MS-SQL database.
type DBClient struct {
	DB *gorm.DB
}

type BackendServiceDBInterface interface {

	//ms-sql
	Connect(connectionString string) (setupError error)
	Ping() error
	Close() error

	BeginTransaction() (txn BackendServiceDBInterface, err *api.APIError)
	CommitTransaction() (err *api.APIError)
	RollbackTransaction() (err *api.APIError)
}
