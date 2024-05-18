package controllers

import (
	service "api/coworkings/application/service"
	"api/coworkings/infrastructure/mappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddImageToCoworkingController(service service.CoworkingsService) gin.HandlerFunc {
	return func(context *gin.Context) {

		requestResult := mappers.AddImageToCoworkingMapper(context)
		if requestResult.IsErr() {
			context.JSON(http.StatusBadRequest, gin.H{"error": requestResult.Err.Error()})
			return
		}

		responseResult := service.AddImageToCoworkingService(requestResult.Ok)
		if responseResult.IsErr() {
			context.JSON(500, gin.H{"message": responseResult.Err.Error()})
			return
		}

		context.JSON(200, gin.H{"response": responseResult.Ok})

	}
}
