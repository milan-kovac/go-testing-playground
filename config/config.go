package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var Env *Environment

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %s", err)
	}

	Env = &Environment{
		Port:       getEnvOrFail("PORT"),
		DBUser:     getEnvOrFail("DB_USER"),
		DBPassword: getEnvOrFail("DB_PASSWORD"),
		DBHost:     getEnvOrFail("DB_HOST"),
		DBPort:     getEnvOrFail("DB_PORT"),
		DBName:     getEnvOrFail("DB_NAME"),
	}
}

func getEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("❌ Missing environment variable %s", key)
	}
	return value
}
