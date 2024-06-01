package controllers

import (
	"exo-planets/api/v1"
	"exo-planets/dataservices"
)

// ControllerDescriber -
type ControllerDescriber interface {
	Connect(connectionString string) (setupError error)
	Ping() error
	Close() error

	BeginTransaction() (txn dataservices.BackendServiceDBInterface, err *api.APIError)
	CommitTransaction() (err *api.APIError)
	RollbackTransaction() (err *api.APIError)
}
