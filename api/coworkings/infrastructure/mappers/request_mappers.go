package mappers

import (
	"api/contracts/requests"
	"api/utils/logger"
	"api/utils/result"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
)

func AddImageToCoworkingMapper(context *gin.Context) result.Result[*requests.AddImageToCoworkingRequest] {

	file, err := context.FormFile("file")

	if err != nil {
		logger.GetLogger().Info("1", err.Error())

		return result.NewResult(&requests.AddImageToCoworkingRequest{}, errors.New("Error opening file"))
	}

	fileContent, err := file.Open()
	if err != nil {
		logger.GetLogger().Info("2", err.Error())

		return result.NewResult(&requests.AddImageToCoworkingRequest{}, errors.New("Error opening file"))
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		logger.GetLogger().Info("2", err.Error())

		return result.NewResult(&requests.AddImageToCoworkingRequest{}, errors.New("Error opening file"))
	}

	coworkingUuid := context.Param("uuid")
	if coworkingUuid == "" {
		return result.NewResult(&requests.AddImageToCoworkingRequest{}, errors.New("Missing coworking uuid"))
	}

	return result.NewResult(&requests.AddImageToCoworkingRequest{
		CoworkingUuid: coworkingUuid,
		ImageData:     fileBytes,
		ImageFilename: file.Filename,
	}, nil)
}


func GetUrlCoworkingUUID(context *gin.Context) result.Result[string] {
	coworkingUuid := context.Param("uuid")
	if coworkingUuid == "" {
		return result.NewResult("", errors.New("Missing coworking uuid"))
	}

	return result.NewResult(coworkingUuid, nil)
}