package dependencies

import "github.com/jinzhu/gorm"

// Depends Store the dependencies
type Depends struct {
	databases DatabaseFactory
}

// GetMainDbConection Get the application database main conection
func (d Depends) GetMainDbConection() *gorm.DB {
	return d.databases.MainConection
}
