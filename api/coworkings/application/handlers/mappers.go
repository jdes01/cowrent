package application

import (
	"api/contracts/requests"
	"api/coworkings/domain"
	r "api/utils/result"
)

func CreateCoworkingRequestMapper(request *requests.CreateCoworkingRequest) r.Result[domain.Coworking] {
	workspacesResult := extractWorkspaces(request.Workspaces)
	if workspacesResult.IsErr() {
		return r.NewResult(domain.Coworking{}, workspacesResult.Err)
	}

	return domain.CreateNewCoworking(request.Name, workspacesResult.Ok)
}

func extractWorkspaces(requests []*requests.CreateWorkspaceRequest) r.Result[[]domain.Workspace] {
	workspaces := make([]domain.Workspace, 0)
	for _, request := range requests {
		workspaceResult := domain.CreateNewWorkspace(request.Name, int(request.Seats))
		if workspaceResult.IsErr() {
			return r.NewResult(make([]domain.Workspace, 0), workspaceResult.Err)
		}
		workspaces = append(workspaces, workspaceResult.Ok)
	}
	return r.NewResult(workspaces, nil)
}
