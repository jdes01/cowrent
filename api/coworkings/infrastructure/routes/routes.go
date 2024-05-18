package controllers

import (
	coworkingService "api/coworkings/application/service"
	controllers "api/coworkings/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureCoworkingsRoutes(router *gin.RouterGroup, service coworkingService.CoworkingsService) {
	router.Group("/coworking").
		POST("", controllers.CreateCoworkingController(service)).
		POST("/:uuid/image", controllers.AddImageToCoworkingController(service))

	router.Group("/coworkings").
		GET("", controllers.GetCoworkingsController(service))
}
