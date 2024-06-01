package main

import (
	"os"

	"gas-giant/config"
	"gas-giant/constants"
	"gas-giant/dataservices"
	"gas-giant/engine"
	"gas-giant/logging"
	"gas-giant/server"

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
