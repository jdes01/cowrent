package application

import (
	"api/contracts/requests"
	domain "api/coworkings/domain"
	r "api/utils/result"

	"github.com/google/uuid"
)

type AddWorkspaceToCoworkingHandler struct {
	Repository   domain.CoworkingRepository
}

func NewAddWorkspaceToCoworkingHandler(repository domain.CoworkingRepository) *AddWorkspaceToCoworkingHandler {
	return &AddWorkspaceToCoworkingHandler{Repository: repository}
}

func (handler *AddWorkspaceToCoworkingHandler) Execute(request *requests.AddWorkspaceToCoworkingRequest) r.Result[domain.WorkspaceUUID] {
	
	coworkingResult := handler.Repository.
		GetCoworkingByUuid(&domain.CoworkingUUID{Value: uuid.MustParse(request.GetCoworkingUuid())}).
		LogThisIfErr("Something went wrong retrieving coworking from repository");
	

	if coworkingResult.IsErr() {
		return r.MapError[domain.Coworking, domain.WorkspaceUUID](coworkingResult)
	}

	workspaceResult := domain.CreateNewWorkspace(request.Name, int(request.Seats))
	if workspaceResult.IsErr() {
		return r.MapError[domain.Workspace, domain.WorkspaceUUID](workspaceResult)
	}

	go func(coworking *domain.Coworking, workspace *domain.Workspace) {
		coworking.AddWorkspace(*workspace)

		handler.Repository.
			SaveCoworking(coworking).
			LogThisIfErr("Error saving the coworking").
			LogThisIfOk("Workspace saved successfully")
		
	}(&coworkingResult.Ok, &workspaceResult.Ok)

	return r.NewResult(workspaceResult.Ok.UUID, nil)
}
