package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment       string // develop, staging, production
	PostgresHost      string
	PostgresPort      int
	PostgresDatasbase  string
	PostgresUser      string
	PostgresPassword  string
	LogLevel          string
	RPCPort           string

	// user service
	UserServiceHost string
	UserServicePort int

	// comment service
	PostServiceHost string
	PostServicePort int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatasbase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "exam"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "123"))

	// user service configuration
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST","user"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9000))

	// comment service configuration
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST","post"))
    c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 8000))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":7000"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

