package application

import (
	"api/contracts/requests"
	domain "api/coworkings/domain"
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

	coworkingResult := handler.Repository.
		GetCoworkingByUuid(&domain.CoworkingUUID{Value: uuid.MustParse(request.GetCoworkingUuid())}).
		LogThisIfErr("Something went wrong retrieving coworking from repository")

	if coworkingResult.IsErr() {
		return r.MapError[domain.Coworking, domain.CoworkingImage](coworkingResult)
	}

	uploaderResult := handler.FileUploader.
		UploadCoworkingImage(&coworkingResult.Ok, request.GetImageFilename(), request.GetImageData()).
		LogThisIfErr("Error uploading image")

	if uploaderResult.IsErr() {
		return r.MapError[string, domain.CoworkingImage](uploaderResult)
	}

	newImage := domain.CoworkingImage{Name: request.ImageFilename, URL: uploaderResult.Ok}

	go func(coworking *domain.Coworking, image *domain.CoworkingImage) {
		coworking.AddImage(*image)

		handler.Repository.
			SaveCoworking(coworking).
			LogThisIfErr("Error saving the coworking")

	}(&coworkingResult.Ok, &newImage)

	return r.NewResult(newImage, nil)

}
