package application

import (
	domain "api/coworkings/domain"
	r "api/utils/result"

	_ "github.com/lib/pq"
)

type GetCoworkingsHandler struct {
	Repository domain.CoworkingRepository
}

func NewGetCoworkingsHandler(repository domain.CoworkingRepository) *GetCoworkingsHandler {
	return &GetCoworkingsHandler{Repository: repository}
}

func (handler *GetCoworkingsHandler) Execute() r.Result[[]domain.Coworking] {
	return handler.Repository.GetCoworkings()
}
