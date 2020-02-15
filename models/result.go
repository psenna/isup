package models

import (
	"github.com/jinzhu/gorm"
	"github.com/psenna/isup/dependencies"
)

// Result A Result from a system test
type Result struct {
	gorm.Model
	System         string
	Test           string
	StatusCode     int
	ResponseTime   float64
	ResponseLength int64
}

// Save Save teh result
func (r Result) Save() {
	dependencies.AppDependencies.GetMainDbConection().Create(r)
}
