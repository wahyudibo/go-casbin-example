package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database                  *DatabaseConfig
	DatabaseMigrationFilePath string `envconfig:"DATABASE_MIGRATION_FILE_PATH"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"DATABASE_HOST"`
	Port     int    `envconfig:"DATABASE_PORT"`
	User     string `envconfig:"DATABASE_USER"`
	Password string `envconfig:"DATABASE_PASSWORD"`
	Name     string `envconfig:"DATABASE_NAME"`
}

// New initializes new config by loading the value from environment variables.
func New() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
