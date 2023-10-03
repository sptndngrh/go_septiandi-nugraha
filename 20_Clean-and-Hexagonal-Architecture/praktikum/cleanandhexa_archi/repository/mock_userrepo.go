package repository

import (
	"cleanandhexa_archi/model"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{}
}

func (m *MockUserRepo) Create(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) Find() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}
