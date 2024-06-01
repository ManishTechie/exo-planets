// Package config  will be basis of environment configs in future
package config

import (
	"gas-giant/logging"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// SetEnvFromDotEnv sets the environment defined in dotenv
func SetEnvFromDotEnv() {
	err := godotenv.Load()
	if err != nil {
		logging.GetLogger().With(zap.Error(err)).Fatal(`Error loading .env file. Copy '.env.sample' to '.env' and add your environment data.
		DO NOT, I repeat DO NOT DELETE .env.sample file. Its part of version control.
		That being said here is the error`)
	}
}

func ensureEnvironment() {
	missingKeys := []string{}
	mapEnv, err := godotenv.Read(".env.sample")
	if err != nil {
		logging.GetLogger().Error("Env sample: ", zap.Error(err))
		return
	}

	for key := range mapEnv {
		_, ok := os.LookupEnv(key)
		if !ok {
			logging.GetLogger().Info("env not found", zap.String("key", key))
			missingKeys = append(missingKeys, key)
		}
	}
	if len(missingKeys) > 0 {
		logging.GetLogger().Panic("Proper environment not set for", zap.Any("missing-keys", missingKeys))
	}
}
