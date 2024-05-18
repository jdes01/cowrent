package tests_handlers

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	handlers "api/coworkings/application/handlers"
	"api/coworkings/domain"
)

func TestGetCoworkingsHandler_Execute(t *testing.T) {
	mockRepo := &mockCoworkingRepository{}
	handler := handlers.NewGetCoworkingsHandler(mockRepo)

	type ExpectedCoworkings struct {
		Coworking     domain.Coworking
		ExpectedError error
	}

	expectedCoworkings := []ExpectedCoworkings{
		{
			Coworking: domain.Coworking{
				UUID:       domain.CoworkingUUID{Value: uuid.New()},
				Name:       domain.CoworkingName{Value: "Coworking A"},
				Workspaces: []domain.Workspace{},
			},
			ExpectedError: nil,
		},
		{
			Coworking: domain.Coworking{
				UUID: domain.CoworkingUUID{Value: uuid.New()},
				Name: domain.CoworkingName{Value: "Coworking B"},
				Workspaces: []domain.Workspace{
					{
						UUID:      domain.WorkspaceUUID{Value: uuid.New()},
						Name:      domain.WorkspaceName{Value: "Workspace C"},
						CreatedAt: time.Now(),
					},
				},
			},
			ExpectedError: nil,
		},
	}

	mockRepo.On("GetCoworkings").Return(func() []domain.Coworking {
		coworkings := make([]domain.Coworking, len(expectedCoworkings))
		for i, expected := range expectedCoworkings {
			coworkings[i] = expected.Coworking
		}
		return coworkings
	}(), nil)

	result := handler.Execute()

	assert.NoError(t, result.Err)
	assert.Len(t, result.Ok, len(expectedCoworkings))

	for i, expected := range expectedCoworkings {
		assert.Equal(t, expected.Coworking.UUID, result.Ok[i].UUID)
		assert.Equal(t, expected.Coworking.Name, result.Ok[i].Name)
		assert.Equal(t, len(expected.Coworking.Workspaces), len(result.Ok[i].Workspaces))

		for j, expectedWorkspace := range expected.Coworking.Workspaces {
			assert.Equal(t, expectedWorkspace, result.Ok[i].Workspaces[j])
		}
	}

	mockRepo.AssertCalled(t, "GetCoworkings")
}
