package config

import (
	"fmt"
	"os"

	"github.com/DamiaoCanndido/gopportunities/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	errEnv := godotenv.Load()
	if errEnv != nil {
		logger.Errorf("postgresql env load error: %v", errEnv)
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbUser,
		dbPass,
		dbDatabase,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgresql opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("postgresql automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
