package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	// user service configuration
	UserServiceHost string
	UserServicePort int

	// product service configuration
	PostServiceHost string
	PostServicePort int

	// comment service configuration
	CommentServiceHost string
	CommentServicePort int

	// context timeout in seconds
	CtxTimeout int


	// casbin
	AuthConfigPath string
	CSVFilePath    string

	// JWT
	SigningKey        string
	AccessTokenTimout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	// user service configuration
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9000))

	// product service configuration
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 8000))

	// comment service configuration
	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "comment"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 7000))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 5))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
