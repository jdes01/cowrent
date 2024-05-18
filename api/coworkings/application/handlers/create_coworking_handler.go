package application

import (
	"api/contracts/requests"
	domain "api/coworkings/domain"
	"api/utils/logger"
	r "api/utils/result"
)

type CreateCoworkingHandler struct {
	Repository domain.CoworkingRepository
}

func NewCreateCoworkingHandler(repository domain.CoworkingRepository) *CreateCoworkingHandler {
	return &CreateCoworkingHandler{Repository: repository}
}

func (handler *CreateCoworkingHandler) Execute(request *requests.CreateCoworkingRequest) r.Result[domain.Coworking] {

	coworkingResult := CreateCoworkingRequestMapper(request)
	if coworkingResult.IsErr() {
		return r.NewResult(domain.Coworking{}, coworkingResult.Err)
	}

	go func(coworking *domain.Coworking) {
		if repositoryResult := handler.Repository.SaveCoworking(coworking); repositoryResult.IsErr() {
			logger.GetLogger().Error("Something went wrong saving user", coworking)
		} else {
			logger.GetLogger().Info("Coworking was saved successfully!", nil)
		}

	}(&coworkingResult.Ok)

	return coworkingResult
}
