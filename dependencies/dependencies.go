package dependencies

import (
	"github.com/jinzhu/gorm"
)

// AppDependencies Export the app dependencies (db, http client, ...) to all packages
var AppDependencies Dependencies

// Dependencies Store the dependencies
type Dependencies struct {
	databases Database
	configs   Configurations
}

// GetMainDbConection Get the application database main conection
func (d Dependencies) GetMainDbConection() *gorm.DB {
	return d.databases.GetMainDataseConnection()
}

// InitDependencies Init the app dependencies
func InitDependencies(configurations Configurations) error {
	AppDependencies = Dependencies{}
	var err error
	AppDependencies.configs = configurations

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
