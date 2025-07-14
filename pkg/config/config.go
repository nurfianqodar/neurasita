package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	DATABASE_URL string
	APP_ID       string
	APP_SECRET   string
	HOST         string
	PORT         int
)

func init() {
	DATABASE_URL = mustGetEnv("DATABASE_URL")
	APP_ID = mustGetEnv("APP_ID")
	APP_SECRET = mustGetEnv("APP_SECRET")
	HOST = getEnv("HOST", "0.0.0.0")

	portString := mustGetEnv("PORT")
	portInt, err := strconv.Atoi(portString)
	if err != nil {
		panic("PORT was not valid number. port must be a number between 0 and 65535")
	}
	PORT = portInt
}

// Helper
func getEnv(key, fb string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fb
}

func mustGetEnv(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	msg := fmt.Sprintf("environment variable %s was not declared", key)
	panic(msg)
}
