package dependencies

import (
	"github.com/jinzhu/gorm"
	"github.com/psenna/isup/models"
)

// AppDependencies Export the app dependencies (db, http client, ...) to all packages
var AppDependencies Dependencies

// Dependencies Store the dependencies
type Dependencies struct {
	databases DatabaseFactory
	configs   models.Config
}

// GetMainDbConection Get the application database main conection
func (d Dependencies) GetMainDbConection() *gorm.DB {
	return d.databases.GetMainDataseConnection()
}

// InitDependencies Init the app dependencies
func InitDependencies() error {
	AppDependencies = Dependencies{}
	var err error
	AppDependencies.configs, err = models.CreateConfig("{}")

	if err != nil {
		return err
	}

	return nil
}

// CloseDependencies Close the initied dependencies
func CloseDependencies() error {
	AppDependencies.databases.Close()
	return nil
}
