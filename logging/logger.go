package logging

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// InitLogger initializes the logger
func InitLogger() {
	once.Do(func() {
		var err error
		// You can customize the logger configuration here
		logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	})
}

// GetLogger returns the initialized logger instance
func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}
