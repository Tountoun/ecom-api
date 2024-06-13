package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	PublicHost string
	Port string
	DBUsername string
	DBPassword string
	DBName string
	DBAddress string
}

var Envs = initConfigurations()

// Get all configurations properties
func initConfigurations() Config {
	godotenv.Load()
	
	return Config{
		PublicHost: getEnv("DB_SERVER_HOST_NAME", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUsername: getEnv("DB_USER", "litt"),
		DBPassword: getEnv("DB_PASS", "litt"),
		DBName: getEnv("DB_NAME", "ecom-api"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
	}
}

// Get environnment variable by using its key
// If not found, return the fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}