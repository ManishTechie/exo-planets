package main

//go:generate mockgen -destination mocks/exoplanet.go -package=mock_dataservices exo-planets/dataservices BackendServiceDBInterface
