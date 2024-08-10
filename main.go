package main

import (
	"os"

	"exo-planets/config"
	"exo-planets/constants"
	"exo-planets/dataservices"
	"exo-planets/engine"
	"exo-planets/logging"
	"exo-planets/server"

	_ "github.com/golang/mock/mockgen/model"

	"go.uber.org/zap"
)

// VERSION will be overwritten by the Go toolchain during release
var VERSION string = "v0.0.1"

func main() {
	logging.GetLogger().Info("Gas Giant")

	config.SetConfigs()

	// set up the database connection
	var db dataservices.BackendServiceDBInterface = &dataservices.DBClient{}
	if dbConnectErr := db.Connect(os.Getenv(constants.DB_CONNECTION_STRING)); dbConnectErr != nil {
		logging.GetLogger().Fatal("Failed to set up dataservices", zap.Error(dbConnectErr))
	}
	db = dataservices.DB()

	// start the server
	server.Start(
		engine.BuildGinEngine(db, VERSION),
		"gas-giant",
		db.Close,
	)

}
