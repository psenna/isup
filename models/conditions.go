package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/psenna/isup/dependencies"
)

// Condition A condition applied to a test result to define the test status
type Condition struct {
	gorm.Model
	Name     string
	Test     Test
	TestID   uint
	Target   string
	Operator string
	Operand  string
	Active   bool
	Valid    bool `gorm:"-"`
}

// Validate Validate a test condition
func (c *Condition) Validate() (valid bool, problems string) {
	c.Valid = true

	return c.Valid, problems
}

// Save Save or create a test condition
func (c Condition) Save() error {
	valid, errs := c.Validate()

	if !valid {
		return errors.New("Invalid Test Condition:\n" + errs)
	}

	return dependencies.AppDependencies.GetMainDbConnection().Save(c).Error
}
