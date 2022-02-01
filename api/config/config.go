package config

import (
	"os"

	"github.com/subosito/gotenv"
)

// DBConnection connection to a database
type DBConnection struct {
	DBDialect         string
	DBConnection      string
	DbMaxIdleConns    int
	DbMaxOpenConns    int
	DbConnMaxLifetime int
	DbLogging         bool
}

type JWTConf struct {
	Secret string
}

// AppConfig application configuration
type AppConfig struct {
	DBConnections map[string]DBConnection
	JWTConf
}

// Load app configuration
func Load() *AppConfig {
	gotenv.Load()

	return &AppConfig{
		DBConnections: map[string]DBConnection{
			"development": {
				DBDialect:         "mysql",
				DBConnection:      os.Getenv("DB_DEV_CONNECTION"),
				DbMaxIdleConns:    10,
				DbMaxOpenConns:    100,
				DbConnMaxLifetime: 30, // minutes
				DbLogging:         true,
			},
		},
		JWTConf: JWTConf{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
