package models

import "github.com/jinzhu/gorm"

// System A system that is monitored
type System struct {
	gorm.Model
	Name     string `json:"name"`
	Group    string `json:"group"`
	Priority uint   `json:"priority"`
	Tests    []Test `json:"tests"`
	Valid    bool   `gorm:"-"`
}

// Initiate Initiate a Test struct
func (s *System) Initiate(sysName string) {

	s.Validate()
}

// Validate Validate
func (s *System) Validate() (valid bool, problems string) {
	s.Valid = true

	if s.Name == "" {
		problems += "A system must have a name\n"
		s.Valid = false
	}

	if len(s.Tests) < 1 {
		problems += "A system must have tests\n"
		s.Valid = false
	}

	return s.Valid, problems
}
