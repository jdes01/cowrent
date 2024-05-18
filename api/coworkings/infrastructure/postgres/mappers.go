package postgres

import (
	domain "api/coworkings/domain"
)

func MapCoworkingToPostgres(coworking domain.Coworking) PostgresCoworking {
	var postgresWorkspaces []PostgresWorkspace
	for _, workspace := range coworking.Workspaces {
		postgresWorkspaces = append(postgresWorkspaces, MapWorkspaceToPostgres(workspace))
	}

	var postgresImages []CoworkingImage
	for _, image := range coworking.Images {
		postgresImages = append(postgresImages, CoworkingImage{
			Name: image.Name,
			URL:  image.URL,
		})
	}

	return PostgresCoworking{
		UUID:       coworking.UUID.Value,
		CreatedAt:  coworking.CreatedAt,
		Name:       coworking.Name.Value,
		Images:     postgresImages,
		Workspaces: postgresWorkspaces,
	}
}

func MapWorkspaceToPostgres(workspace domain.Workspace) PostgresWorkspace {
	return PostgresWorkspace{
		UUID:      workspace.UUID.Value,
		CreatedAt: workspace.CreatedAt,
		Name:      workspace.Name.Value,
		Seats:     workspace.Seats.Value,
	}
}

func MapPostgresCoworkingListToDomain(postgresCoworkings []PostgresCoworking) []domain.Coworking {
	var domainCoworkings []domain.Coworking
	for _, postgresCoworking := range postgresCoworkings {
		domainCoworkings = append(domainCoworkings, MapPostgresCoworkingToDomain(postgresCoworking))
	}

	return domainCoworkings
}

func MapPostgresCoworkingToDomain(pc PostgresCoworking) domain.Coworking {
	var domainWorkspaces []domain.Workspace
	for _, postgresWorkspace := range pc.Workspaces {
		domainWorkspaces = append(domainWorkspaces, MapPostgresWorkspaceToDomain(postgresWorkspace))
	}

	var domainImages []domain.CoworkingImage
	for _, postgresImage := range pc.Images {
		domainImages = append(domainImages, domain.CoworkingImage{
			Name: postgresImage.Name,
			URL:  postgresImage.URL,
		})
	}

	return domain.Coworking{
		UUID:       domain.CoworkingUUID{Value: pc.UUID},
		CreatedAt:  pc.CreatedAt,
		Name:       domain.CoworkingName{Value: pc.Name},
		Images:     domainImages,
		Workspaces: domainWorkspaces,
	}
}

func MapPostgresWorkspaceToDomain(postgresWorkspace PostgresWorkspace) domain.Workspace {
	return domain.Workspace{
		UUID:      domain.WorkspaceUUID{Value: postgresWorkspace.UUID},
		CreatedAt: postgresWorkspace.CreatedAt,
		Name:      domain.WorkspaceName{Value: postgresWorkspace.Name},
		Seats:     domain.WorkspaceSeats{Value: postgresWorkspace.Seats},
	}
}
