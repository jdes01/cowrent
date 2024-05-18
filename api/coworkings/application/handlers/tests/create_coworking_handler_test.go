package tests_handlers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"api/contracts/requests"
	r "api/utils/result"

	application "api/coworkings/application/handlers"
	"api/coworkings/domain"
)

func TestCreateCoworkingHandler_Execute(t *testing.T) {

	type WorkspaceTest struct {
		Workspace    requests.CreateWorkspaceRequest
		ErrorMessage string
	}

	TESTING_WORKSPACE_REQUEST_LIST := []WorkspaceTest{
		{
			Workspace: requests.CreateWorkspaceRequest{
				Name:  "",
				Seats: 5,
			},
			ErrorMessage: "Workspace name cannot be empty",
		},
		{
			Workspace: requests.CreateWorkspaceRequest{
				Name:  "Workspace 2",
				Seats: 0,
			},
			ErrorMessage: "Workspace gotta have seats",
		},
		{
			Workspace: requests.CreateWorkspaceRequest{
				Name:  "Workspace 2",
				Seats: -1,
			},
			ErrorMessage: "Workspace gotta have seats",
		},
	}

	EMPTY_WORKSPACE_LIST := make([]domain.Workspace, 0)
	EMPTY_WORKSPACE_REQUEST_LIST := make([]*requests.CreateWorkspaceRequest, 0)

	type TestArgs struct {
		TestName         string
		Request          requests.CreateCoworkingRequest
		RepositoryReturn r.Result[domain.Coworking]
		ExpectedResponse r.Result[domain.Coworking]
	}

	testCases := []TestArgs{
		{
			TestName:         "Empty coworking name",
			Request:          requests.CreateCoworkingRequest{Name: "", Workspaces: EMPTY_WORKSPACE_REQUEST_LIST},
			RepositoryReturn: r.NewResult(domain.Coworking{}, nil),
			ExpectedResponse: r.NewResult(domain.NewCoworking(domain.CoworkingName{}, EMPTY_WORKSPACE_LIST), errors.New("Coworking name cannot be empty")),
		},
		{
			TestName:         "Create coworking successfully",
			Request:          requests.CreateCoworkingRequest{Name: "Sample Coworking 1", Workspaces: EMPTY_WORKSPACE_REQUEST_LIST},
			RepositoryReturn: r.NewResult(domain.Coworking{}, nil),
			ExpectedResponse: r.NewResult(domain.NewCoworking(domain.CoworkingName{Value: "Sample Coworking 1"}, EMPTY_WORKSPACE_LIST), nil),
		},
	}

	for _, testing_workspace_request := range TESTING_WORKSPACE_REQUEST_LIST {
		testCases = append(testCases, TestArgs{
			TestName:         "Possible coworking list",
			Request:          requests.CreateCoworkingRequest{Name: "Sample Coworking 1", Workspaces: []*requests.CreateWorkspaceRequest{&testing_workspace_request.Workspace}},
			RepositoryReturn: r.NewResult(domain.Coworking{}, nil),
			ExpectedResponse: r.NewResult(domain.NewCoworking(domain.CoworkingName{}, EMPTY_WORKSPACE_LIST), errors.New(testing_workspace_request.ErrorMessage)),
		})
	}

	mockRepo := &mockCoworkingRepository{}
	handler := application.NewCreateCoworkingHandler(mockRepo)

	for _, testCase := range testCases {

		mockRepo.On("SaveCoworking", mock.AnythingOfType("*domain.Coworking")).Return(testCase.RepositoryReturn, nil).Once()

		t.Run(testCase.TestName, func(t *testing.T) {

			switch response := handler.Execute(&testCase.Request); {
			case response.IsOk():
				assert.Equal(t, testCase.ExpectedResponse.Ok.Name, response.Ok.Name)
				assert.Equal(t, len(testCase.ExpectedResponse.Ok.Workspaces), len(response.Ok.Workspaces))
			case response.IsErr():
				assert.Equal(t, testCase.ExpectedResponse.Err.Error(), response.Err.Error())
			}

		})
	}
}
