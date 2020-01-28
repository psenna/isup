package dependencies

// Configurations The app configurations
type Configurations struct {
	DatabaseType string `json:"database_type"`
	DatabaseURL  string `json:"database_url"`
}

// GetDatabaseConnection Get database type and url from connection
func (c Configurations) GetDatabaseConnection() (databaseType string, databaseURL string) {
	return c.DatabaseType, c.DatabaseURL
}
