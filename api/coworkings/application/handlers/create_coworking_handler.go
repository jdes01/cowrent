package application

import (
	"api/contracts/requests"
	domain "api/coworkings/domain"
	r "api/utils/result"
)

type CreateCoworkingHandler struct {
	Repository domain.CoworkingRepository
}

func NewCreateCoworkingHandler(repository domain.CoworkingRepository) *CreateCoworkingHandler {
	return &CreateCoworkingHandler{Repository: repository}
}

func (handler *CreateCoworkingHandler) Execute(request *requests.CreateCoworkingRequest) r.Result[domain.Coworking] {

	if coworkingResult := CreateCoworkingRequestMapper(request); coworkingResult.IsErr() {
		return coworkingResult

	} else {
		go func(coworking *domain.Coworking) {
			handler.Repository.
				SaveCoworking(coworking).
				LogThisIfErr("Something went wrong saving user", coworking).
				LogThisIfOk("Coworking was saved successfully!")

		}(&coworkingResult.Ok)

		return coworkingResult
	}
}
