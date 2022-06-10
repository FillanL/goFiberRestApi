package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Environment struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	PORT string
	TempSecret string
	DBbUrl string
}

var Env *Environment

func getEnv(key string, required bool) string {
	value, ok := os.LookupEnv(key)
	if !ok && required {
		log.Fatalf("Missing or invalid environment key: '%s'", key)
	}
	return value
}

func LoadEnvironment() {
	if Env == nil {
		Env = new(Environment)
	}
	Env.DBHost = getEnv("DB_HOST", false)
	Env.DBPort = getEnv("DB_PORT", false)
	Env.DBUser = getEnv("DB_USER", false)
	Env.DBPass = getEnv("DB_PASS", false)
	Env.DBName = getEnv("DB_NAME", false)
	Env.DBbUrl = getEnv("DATABASE_URL", false)
	Env.PORT = getEnv("PORT", false)
	Env.TempSecret = getEnv("TempSecret", false)
}

func LoadEnvironmentFile() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error on load environment")
	}
	LoadEnvironment()
}