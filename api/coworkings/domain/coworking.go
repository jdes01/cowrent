package domain

import (
	r "api/utils/result"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CoworkingUUID struct {
	Value uuid.UUID
}

type CoworkingName struct {
	Value string
}

func CoworkingNameFromString(value string) r.Result[CoworkingName] {
	if value == "" {
		return r.NewResult(CoworkingName{}, errors.New("Coworking name cannot be empty"))
	}
	return r.NewResult(CoworkingName{Value: value}, nil)
}

type CoworkingImage struct {
	Name string
	URL  string
}

type Coworking struct {
	UUID       CoworkingUUID
	CreatedAt  time.Time
	Name       CoworkingName
	Images     []CoworkingImage
	Workspaces []Workspace
}

func NewCoworking(name CoworkingName, workspaces []Workspace) Coworking {
	return Coworking{
		UUID:       CoworkingUUID{Value: uuid.New()},
		CreatedAt:  time.Now(),
		Name:       name,
		Workspaces: workspaces,
	}
}

func CreateNewCoworking(name string, workspaces []Workspace) r.Result[Coworking] {

	nameResult := CoworkingNameFromString(name)
	if nameResult.IsErr() {
		return r.NewResult(Coworking{}, nameResult.Err)
	}

	return r.NewResult(NewCoworking(nameResult.Ok, workspaces), nil)
}

func (coworking *Coworking) AddImage(image CoworkingImage) {
	coworking.Images = append(coworking.Images, image)
}

func (coworking *Coworking) AddWorkspace(workspace Workspace) {
	coworking.Workspaces = append(coworking.Workspaces, workspace)
}