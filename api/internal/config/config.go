package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBTz       string
}

func loadDotEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		_ = godotenv.Load(".env")
	}
}

func getEnv(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("%s is not set", key)
	}
	return v, nil
}

func Load() (Config, error) {
	loadDotEnv()

	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return Config{}, err
	}

	dbUser, err := getEnv("DB_USER")
	if err != nil {
		return Config{}, err
	}

	dbPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		return Config{}, err
	}

	dbName, err := getEnv("DB_NAME")
	if err != nil {
		return Config{}, err
	}

	dbPort, err := getEnv("DB_PORT")
	if err != nil {
		return Config{}, err
	}

	dbSSlMode, err := getEnv("DB_SSLMODE")
	if err != nil {
		return Config{}, err
	}

	dbTz, err := getEnv("DB_TZ")
	if err != nil {
		return Config{}, err
	}

	return Config{
		DBHost:     dbHost,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBPort:     dbPort,
		DBSSLMode:  dbSSlMode,
		DBTz:       dbTz,
	}, nil
}
