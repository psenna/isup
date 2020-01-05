package dependencies

import (
	"github.com/jinzhu/gorm"
	"log"

	// Import portgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// Import sqlite dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DatabaseFactory Create database conections
type DatabaseFactory struct {
	mainConnection *gorm.DB
}

// GetMainDataseConnection Get main app db connection
func (d *DatabaseFactory) GetMainDataseConnection() *gorm.DB {
	if d.mainConnection == nil {
		db, err := d.createMainDatabaseConnection()
		if err == nil {
			d.mainConnection = db
		} else {
			log.Println(err.Error())
		}
	}

	return d.mainConnection
}

func (d DatabaseFactory) createMainDatabaseConnection() (*gorm.DB, error) {
	return d.createDatabaseConnection(AppDependencies.configs.GetDatabaseConnection())
}

func (d DatabaseFactory) createDatabaseConnection(dbType string, url string) (*gorm.DB, error) {
	return gorm.Open(dbType, url)
}

// Close Close the databases connections
func (d *DatabaseFactory) Close() {
	if d.mainConnection != nil {
		d.mainConnection.Close()
	}
}
