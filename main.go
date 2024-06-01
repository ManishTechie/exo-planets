package main

import (
	"gas-giant/config"
	"gas-giant/logging"
)

// VERSION will be overwritten by the Go toolchain during release
var VERSION string = "v0.0.1"

func main() {
	logging.GetLogger().Info("Gas Giant")

	config.SetConfigs()

}
