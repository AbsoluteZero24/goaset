package database

import (
	"fmt"
	"log"

	"github.com/AbsoluteZero24/goaset/internal/config"
	"github.com/AbsoluteZero24/goaset/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize(dbConfig config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Jakarta",
		dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	for _, model := range models.RegisterModels() {
		err := db.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database Migrated Successfully")
}
