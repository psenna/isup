package dependencies

import (
	"github.com/jinzhu/gorm"

	// Import portgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// Import sqlite dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DatabaseFactory Create database conections
type DatabaseFactory struct {
	MainConection *gorm.DB
}
