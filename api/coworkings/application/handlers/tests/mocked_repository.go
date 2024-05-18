package tests_handlers

import (
	"api/coworkings/domain"
	r "api/utils/result"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockCoworkingRepository struct {
	mock.Mock
}

func (m *mockCoworkingRepository) SaveCoworking(coworking *domain.Coworking) r.Result[domain.Coworking] {
	args := m.Called(coworking)
	return r.NewResult(args.Get(0).(domain.Coworking), args.Error(1))
}

func (m *mockCoworkingRepository) GetCoworkings() r.Result[[]domain.Coworking] {
	args := m.Called()
	return r.NewResult(args.Get(0).([]domain.Coworking), args.Error(1))
}

func (m *mockCoworkingRepository) GetCoworkingByUuid(uuid *domain.CoworkingUUID) r.Result[domain.Coworking] {
	args := m.Called()
	return r.NewResult(args.Get(0).(domain.Coworking), args.Error(1))
}

func (m *mockCoworkingRepository) AssertExpectations(t *testing.T) {
	m.Mock.AssertExpectations(t)
}

func (m *mockCoworkingRepository) GetCoworkingsWithWorkspaces() r.Result[[]domain.Coworking] {
	args := m.Called()
	return r.NewResult(args.Get(0).([]domain.Coworking), args.Error(1))
}
