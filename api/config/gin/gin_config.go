package gin_config

import (
	"strconv"

	"github.com/eldimious/golang-api-showcase/utils/env"
	"github.com/gin-gonic/gin"
)

func GetGinModeFromEnv(gin_mode string) string {

	switch gin_mode {
	case "debug":
		return gin.DebugMode
	case "release":
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}

type Config struct {
	GinMode        string
	Port           string
	GinLogsEnabled bool
}

func NewConfig() *Config {

	env.CheckDotEnv()

	var gin_mode string

	if mode := env.MustGet("ENV"); mode == "release" {
		gin_mode = gin.ReleaseMode

	} else {
		gin_mode = gin.DebugMode
	}

	port := env.MustGet("PORT")
	if port == "" {
		port = "8080"
	}

	gin_logs_enabled, err := strconv.ParseBool(env.MustGet("GIN_LOGS_ENALBED"))
	if err != nil {
		gin_logs_enabled = false
	}

	return &Config{
		GinMode:        gin_mode,
		Port:           port,
		GinLogsEnabled: gin_logs_enabled,
	}
}
