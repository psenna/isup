package models

import (
	"encoding/json"
	"errors"
)

// Config The app config
type Config struct {
	DatabaseType string   `json:"database_type"`
	DatabaseURL  string   `json:"database_url"`
	Systems      []System `json:"systems"`
}

// Validate Return the configuratons validation erros
func (c Config) Validate() error {
	if len(c.Systems) < 1 {
		return errors.New("At least one configured system required")
	}
	return nil
}

// CreateConfig Create a config object
func CreateConfig(jsonConfigs string) (Config, error) {
	configuration := Config{
		DatabaseType: "sqlite",
		DatabaseURL:  "./Config/gorm.db",
	}

	err := json.Unmarshal([]byte(jsonConfigs), &configuration)

	if err != nil {
		return configuration, errors.New("Invalid json")
	}

	return configuration, configuration.Validate()
}

// GetDatabaseConnection Get database type and url from connection
func (c Config) GetDatabaseConnection() (databaseType string, databaseURL string) {
	return c.DatabaseType, c.DatabaseURL
}
