package controllers

import (
	"api/contracts/requests"
	service "api/coworkings/application/service"
	"api/coworkings/infrastructure/mappers"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddWorkspaceToCoworkingController(service service.CoworkingsService) gin.HandlerFunc {
	return func(context *gin.Context) {

		urlCoworkingUUIDResult := mappers.GetUrlCoworkingUUID(context)
		if urlCoworkingUUIDResult.IsErr() {
			context.JSON(http.StatusBadRequest, gin.H{"error": urlCoworkingUUIDResult.Err.Error()})
			return
		}
		
		requestResult := utils.MapRequestBody[requests.AddWorkspaceToCoworkingRequest](context)
		if requestResult.IsErr() {
			context.JSON(http.StatusBadRequest, gin.H{"error": requestResult.Err.Error()})
			return
		}

		requestResult.Ok.CoworkingUuid = urlCoworkingUUIDResult.Ok

		responseResult := service.AddWorkspaceToCoworkingService(&requestResult.Ok)
		if responseResult.IsErr() {
			context.JSON(500, gin.H{"message": responseResult.Err.Error()})
			return
		}

		context.JSON(200, gin.H{"response": responseResult.Ok})

	}
}
