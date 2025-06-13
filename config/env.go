package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string

	MaxMindUserId     string
	MaxMindLicenseKey string
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:        getEnv("PUBLIC_HOST", "localhost"),
		Port:              getEnv("PORT", "8080"),
		DBUser:            getEnv("DB_USER", "root"),
		DBPassword:        getEnv("DB_PASSWORD", "password"),
		DBAddress:         fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:            getEnv("DB_NAME", "mydb"),
		MaxMindUserId:     getEnv("MAXMIND_USER_ID", ""),
		MaxMindLicenseKey: getEnv("MAXMIND_LICENSE_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
