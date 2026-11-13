package db

import (
	"fmt"
	"log"

	"github.com/Jet1a/go-try-on/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) (*gorm.DB){

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
		cfg.DBTz,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := Migrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed!")

	log.Println("Successfully connected to PostgreSQL!")

	DB = db

	return DB
}
