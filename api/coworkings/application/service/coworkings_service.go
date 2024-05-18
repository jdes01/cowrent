package application

import (
	"api/contracts/dtos"
	"api/contracts/requests"
	"api/contracts/responses"
	handlers "api/coworkings/application/handlers"
	"api/coworkings/domain"
	"api/coworkings/infrastructure/mappers"
	r "api/utils/result"
)

type CoworkingsService struct {
	AddImageToCoworkingHandler handlers.AddImageToCoworkingHandler
	CreateCoworkingHandler     handlers.CreateCoworkingHandler
	GetCoworkingsHandler       handlers.GetCoworkingsHandler
}

func New(repository domain.CoworkingRepository, fileUploader domain.FileUploader) *CoworkingsService {
	return &CoworkingsService{
		AddImageToCoworkingHandler: *handlers.NewAddImageToCoworkingHandler(repository, fileUploader),
		CreateCoworkingHandler:     *handlers.NewCreateCoworkingHandler(repository),
		GetCoworkingsHandler:       *handlers.NewGetCoworkingsHandler(repository),
	}
}

func (service *CoworkingsService) CreateCoworking(request *requests.CreateCoworkingRequest) r.Result[responses.CreateCoworkingResponse] {
	switch coworkingCreatedResult := service.CreateCoworkingHandler.Execute(request); {
	case coworkingCreatedResult.IsErr():
		return r.NewResult(responses.CreateCoworkingResponse{}, coworkingCreatedResult.Err)
	default:
		return r.NewResult(responses.CreateCoworkingResponse{UUID: coworkingCreatedResult.Ok.UUID.Value.String()}, nil)
	}
}

func (service *CoworkingsService) GetCoworkings() r.Result[[]dtos.CoworkingDTO] {
	switch coworkingsResult := service.GetCoworkingsHandler.Execute(); {
	case coworkingsResult.IsErr():
		return r.NewResult([]dtos.CoworkingDTO{}, coworkingsResult.Err)
	default:
		return r.NewResult(mappers.MapCoworkingListToDTO(coworkingsResult.Ok), nil)
	}
}

func (service *CoworkingsService) AddImageToCoworkingService(request *requests.AddImageToCoworkingRequest) r.Result[responses.AddImageToCoworkingResponse] {
	switch newImageResult := service.AddImageToCoworkingHandler.Execute(request); {
	case newImageResult.IsErr():
		return r.NewResult(responses.AddImageToCoworkingResponse{}, newImageResult.Err)
	default:
		return r.NewResult(responses.AddImageToCoworkingResponse{ImagePath: newImageResult.Ok.URL}, nil)
	}
}
