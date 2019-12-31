package dependencies

import "github.com/jinzhu/gorm"

import "github.com/psenna/isup/models"

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
func InitDependencies() {
	AppDependencies = Dependencies{}
}

// CloseDependencies Close de initied dependencies
func CloseDependencies() {
	AppDependencies.databases.Close()
}
