package dataservices

import (
	"sync"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	// appDB is the application's database connection
	appDB *DBClient
	// setupOnce ensures that the connection can be setup only once
	setupOnce sync.Once
)

// Connect sets up the global database connection with sensible defaults.
func (ms *DBClient) Connect(connectionString string) (setupError error) {
	setupOnce.Do(func() {

		var migrations []*gormigrate.Migration
		// If APP_DB_AUTO_MIGRATE is set to true, set the migrations,
		// else leave it on nil, so no migration will take place

		db, err := NewPostgreSQLGormDB(
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			},
			connectionString,
			20,
			3,
			time.Second*30,
			migrations,
		)

		if err != nil {
			setupError = errors.Wrap(err, "Failed to connect")
		}
		appDB = &DBClient{DB: db}

	})

	return
}

// Close the connection to the database.
func (ms *DBClient) Close() error {
	sqlDB, err := ms.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Ping the database.
func (ms *DBClient) Ping() error {
	sqlDB, err := ms.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// DB returns the global database connection.
func DB() *DBClient {
	return appDB
}
