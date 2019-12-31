package models_test

import (
	"errors"
	"testing"

	"github.com/psenna/isup/models"
	"github.com/stretchr/testify/assert"
)

func TestConfigInitialization(t *testing.T) {
	var tests = []struct {
		configJson string
		config     models.Config
		err        error
	}{
		{"{}", models.Config{DatabaseType: "sqlite", DatabaseURL: "./Config/gorm.db"}, errors.New("At least one configured system required")},
		{"{}", models.Config{DatabaseType: "sqlite", DatabaseURL: "./Config/gorm.db"}, errors.New("At least one configured system required")},
		{"{\"database_type\":\"postgres\"}", models.Config{DatabaseType: "postgres", DatabaseURL: "./Config/gorm.db"}, errors.New("At least one configured system required")},
		{"{\"database_url:\"localhost:6154\"}", models.Config{DatabaseType: "sqlite", DatabaseURL: "./Config/gorm.db"}, errors.New("Invalid json")},
		{"{\"systems\":[{\"name\":\"sis1\"}]}", models.Config{DatabaseType: "sqlite", DatabaseURL: "./Config/gorm.db", Systems: []models.System{models.System{Name: "sis1"}}}, nil},
	}

	for _, test := range tests {
		config, err := models.CreateConfig(test.configJson)

		assert.NotNil(t, config)

		assert.True(t, assert.ObjectsAreEqualValues(test.config, config), "The config object are different from the expected")

		if test.err == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.err.Error(), err.Error())
		}
	}
}

func TestConfigGetDatabaseConnection(t *testing.T) {
	var tests = []struct {
		config       models.Config
		databaseType string
		databaseURL  string
	}{
		{models.Config{DatabaseType: "sqlite", DatabaseURL: "./Config/gorm.db"}, "sqlite", "./Config/gorm.db"},
		{models.Config{DatabaseType: "postgres", DatabaseURL: "./Config/gorm.db"}, "postgres", "./Config/gorm.db"},
		{models.Config{DatabaseType: "sqlite", DatabaseURL: "localhost:6565"}, "sqlite", "localhost:6565"},
	}

	for _, test := range tests {
		databaseType, databaseURL := test.config.GetDatabaseConnection()

		assert.Equal(t, test.databaseType, databaseType)
		assert.Equal(t, test.databaseURL, databaseURL)
	}
}
