package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	PublicHost string
	Port string
	DBUsername string
	DBPassword string
	DBName string
	DBAddress string
	JWTExpirationTimeInSeconds int64
	JWTSecret string
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
		JWTExpirationTimeInSeconds: getEnvValueAsInt("JWT_EXP", 3600 * 24 * 7),
		JWTSecret: getEnv("JWT_SECRET", "here-jwt-secret-and-not-yours"),
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

// Get environnment variable as int
func getEnvValueAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return intVal
	}
	return fallback
}