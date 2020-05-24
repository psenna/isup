package models

import (
	"github.com/jinzhu/gorm"
	"github.com/psenna/isup/dependencies"
)

// Result A Result from a system test
type Result struct {
	gorm.Model
	System         System
	SystemID       uint
	Test           string
	StatusCode     int
	ResponseTime   float64
	ResponseLength int64
	Valid          bool `gorm:"-"`
}

// Save Save the result
func (r Result) Save() {
	dependencies.AppDependencies.GetMainDbConnection().Create(r)
}
