package models

import "github.com/jinzhu/gorm"

// Result A Result from a system test
type Result struct {
	gorm.Model
	System         string
	Test           string
	StatusCode     int
	ResponseTime   float64
	ResponseLength int64
}
