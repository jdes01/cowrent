package domain

import (
	r "api/utils/result"
	"errors"
	"time"

	"github.com/google/uuid"
)

type WorkspaceUUID struct {
	Value uuid.UUID
}

type WorkspaceName struct {
	Value string
}

func WorkspaceNameFromString(value string) r.Result[WorkspaceName] {
	if value == "" {
		return r.NewResult(WorkspaceName{}, errors.New("Workspace name cannot be empty"))
	}
	return r.NewResult(WorkspaceName{Value: value}, nil)
}

type WorkspaceSeats struct {
	Value int
}

func WorkspaceSeatsFromInt(value int) r.Result[WorkspaceSeats] {
	if value <= 0 {
		return r.NewResult(WorkspaceSeats{}, errors.New("Workspace gotta have seats"))
	}
	return r.NewResult(WorkspaceSeats{Value: value}, nil)
}

type Workspace struct {
	UUID          WorkspaceUUID
	CoworkingUUID CoworkingUUID
	CreatedAt     time.Time
	Name          WorkspaceName
	Seats         WorkspaceSeats
}

func NewWorkspace(name WorkspaceName, seats WorkspaceSeats) Workspace {
	return Workspace{
		UUID:      WorkspaceUUID{Value: uuid.New()},
		CreatedAt: time.Now(),
		Name:      name,
		Seats:     seats,
	}
}

func CreateNewWorkspace(name string, seats int) r.Result[Workspace] {

	nameResult := WorkspaceNameFromString(name)
	if err := nameResult.Err; err != nil {
		return r.NewResult(Workspace{}, err)
	}

	seatsResult := WorkspaceSeatsFromInt(seats)
	if err := seatsResult.Err; err != nil {
		return r.NewResult(Workspace{}, err)
	}

	return r.NewResult(NewWorkspace(nameResult.Ok, seatsResult.Ok), nil)
}
