package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Load environment variables
	dbConfig := LoadEnv()

	dsn := "host=" + dbConfig.Host +
		" user=" + dbConfig.Username +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.DBName +
		" port=" + dbConfig.Port +
		" sslmode=" + dbConfig.SSLMode

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
