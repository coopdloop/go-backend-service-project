package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port          string
	PolygonApiKey string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:          getEnv("PORT", "3000"),
		PolygonApiKey: getEnvOrError("POLYGON_API_KEY"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvOrError(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s is not set", key))
}
