package router

import (
	"net/http"

	logger_middleware "api/config/middleware/logger"
	prometheus_middleware "api/config/middleware/prometheus"
	coworkingsService "api/coworkings/application/service"
	coworkingRoutes "api/coworkings/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(coworkingsService coworkingsService.CoworkingsService, ginLogsEnabled bool) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	if ginLogsEnabled {
		router.Use(gin.Logger())
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	router.Use(cors.New(config))
	router.Use(logger_middleware.LoggerMiddleware())
	router.Use(prometheus_middleware.PrometheusMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	coworkingRoutes.ConfigureCoworkingsRoutes(router.Group("/api"), coworkingsService)

	return router
}
