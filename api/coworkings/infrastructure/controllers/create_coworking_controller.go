package controllers

import (
	"api/contracts/requests"
	service "api/coworkings/application/service"
	"api/utils"

	"github.com/gin-gonic/gin"
)

func CreateCoworkingController(service service.CoworkingsService) gin.HandlerFunc {
	return func(context *gin.Context) {

		request := utils.MapRequestBody[requests.CreateCoworkingRequest](context)
		if request.IsErr() {
			context.JSON(400, gin.H{"message": "Wrong format"})
			return
		}

		responseResult := service.CreateCoworking(&request.Ok)
		if responseResult.IsErr() {
			context.JSON(500, gin.H{"message": responseResult.Err.Error()})
			return
		}

		context.JSON(200, gin.H{"response": responseResult.Ok})
	}
}
