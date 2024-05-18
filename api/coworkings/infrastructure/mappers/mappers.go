package mappers

import (
	"api/contracts/dtos"
	"api/coworkings/domain"
)

func MapCoworkingListToDTO(domainCoworkings []domain.Coworking) []dtos.CoworkingDTO {
	var coworkingResponses []dtos.CoworkingDTO
	for _, coworking := range domainCoworkings {
		coworkingResponse := MapCoworkingToDTO(coworking)
		coworkingResponses = append(coworkingResponses, coworkingResponse)
	}
	return coworkingResponses
}

func MapCoworkingToDTO(domainCoworking domain.Coworking) dtos.CoworkingDTO {
	workspaceResponses := MapWorkspacesToDTO(domainCoworking.Workspaces)
	imagePaths := make([]string, len(domainCoworking.Images))
	for i, image := range domainCoworking.Images {
		imagePaths[i] = image.URL
	}
	coworkingResponse := dtos.CoworkingDTO{
		UUID:       domainCoworking.UUID.Value.String(),
		Name:       domainCoworking.Name.Value,
		ImagePath:  imagePaths,
		Workspaces: workspaceResponses,
	}
	return coworkingResponse
}

func MapWorkspacesToDTO(domainWorkspaces []domain.Workspace) []*dtos.WorkspaceDTO {
	workspaceResponses := make([]*dtos.WorkspaceDTO, 0)
	for _, domainWorkspace := range domainWorkspaces {
		workspaceResponses = append(workspaceResponses, MapWorkspaceToDTO(domainWorkspace))
	}
	return workspaceResponses
}

func MapWorkspaceToDTO(domainWorkspace domain.Workspace) *dtos.WorkspaceDTO {
	return &dtos.WorkspaceDTO{
		UUID:  domainWorkspace.UUID.Value.String(),
		Name:  domainWorkspace.Name.Value,
		Seats: int32(domainWorkspace.Seats.Value),
	}
}
