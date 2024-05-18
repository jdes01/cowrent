package main

import (
	aws_config "api/config/aws"
	gin_config "api/config/gin"
	postgres_config "api/config/postgres"
	coworkingsService "api/coworkings/application/service"
	aws "api/coworkings/infrastructure/aws"
	coworkingsPostgresRepository "api/coworkings/infrastructure/postgres"
	"api/router"
	"api/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.InitializeLogger("API", true)

	logger.Info("Loading config...", nil)
	postgresConfig := postgres_config.NewConfig()
	awsConfig := aws_config.NewConfig()
	ginConfig := gin_config.NewConfig()

	logger.Info("Setting config config environment", ginConfig.GinMode)
	gin.SetMode(ginConfig.GinMode)

	logger.Info("Loading use cases...", nil)
	coworkingsRepository := coworkingsPostgresRepository.New(postgresConfig.DB)
	fileUploader := aws.NewFileUploader(awsConfig.S3_service)
	coworkingsService := *coworkingsService.New(coworkingsRepository, fileUploader)

	logger.Info("Connecting router...", nil)
	router := router.NewRouter(coworkingsService, ginConfig.GinLogsEnabled)

	logger.Info("Ready to go!!", nil)
	if err := http.ListenAndServe(":"+ginConfig.Port, router); err != nil {
		panic(err)
	}

}
