package application

import (
	"api/contracts/requests"
	domain "api/coworkings/domain"
	"api/utils/logger"
	r "api/utils/result"

	"github.com/google/uuid"
)

type AddImageToCoworkingHandler struct {
	Repository   domain.CoworkingRepository
	FileUploader domain.FileUploader
}

func NewAddImageToCoworkingHandler(repository domain.CoworkingRepository, fileUploader domain.FileUploader) *AddImageToCoworkingHandler {
	return &AddImageToCoworkingHandler{Repository: repository, FileUploader: fileUploader}
}

func (handler *AddImageToCoworkingHandler) Execute(request *requests.AddImageToCoworkingRequest) r.Result[domain.CoworkingImage] {

	coworkingResult := handler.Repository.GetCoworkingByUuid(&domain.CoworkingUUID{Value: uuid.MustParse(request.GetCoworkingUuid())})
	if coworkingResult.IsErr() {
		logger.GetLogger().Error("Something went wrong retrieving coworking from repository", coworkingResult.Err)
		return r.NewResult(domain.CoworkingImage{}, coworkingResult.Err)
	}

	coworking := coworkingResult.Ok

	uploaderResult := handler.FileUploader.UploadCoworkingImage(&coworking, request.GetImageFilename(), request.GetImageData())
	if uploaderResult.IsErr() {
		logger.GetLogger().Error("Error uploading image", uploaderResult.Err)
		return r.NewResult(domain.CoworkingImage{}, uploaderResult.Err)
	}

	newImage := domain.CoworkingImage{Name: request.ImageFilename, URL: uploaderResult.Ok}
	coworking.AddImage(newImage)

	repositoryResult := handler.Repository.SaveCoworking(&coworking)
	if repositoryResult.IsErr() {
		logger.GetLogger().Error("Error saving the coworking", repositoryResult.Err)
		return r.NewResult(domain.CoworkingImage{}, repositoryResult.Err)
	}

	return r.NewResult(newImage, nil)
}
