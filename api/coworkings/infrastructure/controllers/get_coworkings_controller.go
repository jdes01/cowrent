package controllers

import (
	service "api/coworkings/application/service"

	"github.com/gin-gonic/gin"
)

func GetCoworkingsController(service service.CoworkingsService) gin.HandlerFunc {
	return func(context *gin.Context) {

		switch coworkingDTOResult := service.GetCoworkings(); {
		case coworkingDTOResult.IsErr():
			context.JSON(500, gin.H{"message": coworkingDTOResult.Err})

		case coworkingDTOResult.IsOk():
			context.JSON(200, gin.H{"coworkings": coworkingDTOResult.Ok})
		}
	}
}
