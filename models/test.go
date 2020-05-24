package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Test A test in a system
type Test struct {
	gorm.Model
	TestType       string                 `json:"test_type"`
	URL            string                 `json:"url"`
	Headers        map[string]interface{} `json:"headers"`
	FormParameters map[string]interface{} `json:"form_parameters"`
	Interval       int                    `json:"interval"`
	Name           string                 `json:"name"`
	System         System
	SystemID       uint
	LastRun        time.Time
	Valid          bool `gorm:"-"`
}

var validTests = map[string]bool{
	"HTTP-GET":  true,
	"HTTP-POST": true,
}

// Initiate Initiate a Test struct
func (t *Test) Initiate(sysName string) {
	t.System = sysName
	if t.Interval < 0 {
		t.Interval = 60
	}

	t.Validate()
}

// Validate Validate a test
func (t *Test) Validate() (valid bool, problems string) {
	t.Valid = true

	if t.Name == "" {
		problems += "A test must have a name\n"
		t.Valid = false
	}

	if _, ok := validTests[strings.ToUpper(t.TestType)]; !ok {
		problems += "The type " + t.TestType + " is invalid\n"
		t.Valid = false
	}

	if t.TestType == "" {
		problems += "A test must have a type\n"
		t.Valid = false
	}

	if t.URL == "" {
		problems += "A test must have a url\n"
		t.Valid = false
	}

	if t.System == "" {
		problems += "A test must have a system\n"
		t.Valid = false
	}

	return t.Valid, problems
}

// Run Run the test
func (t *Test) Run() Result {
	return Result{}
}

// Run Run the test
func (t Test) runHTTPTest() Result {
	return Result{}
}
