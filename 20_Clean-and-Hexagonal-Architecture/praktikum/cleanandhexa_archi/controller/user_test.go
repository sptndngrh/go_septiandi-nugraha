package controller

import (
	"cleanandhexa_archi/model"
	"cleanandhexa_archi/repository"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteUsers struct {
	suite.Suite
	control *userController
	mock    *repository.MockUserRepo
}

func (s *suiteUsers) SetupSuite() {
	mock := &repository.MockUserRepo{}
	s.mock = mock

	s.control = &userController{
		userRepo: s.mock,
	}
}

func (s *suiteUsers) TestGetAllUsers() {
	s.mock.On("Find").Return([]model.User{
		{
			Email:    "admin@gmail.com",
			Password: "admin123",
		},
	}, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       model.User
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			true,
			model.User{
				Email:    "admin@gmail.com",
				Password: "admin123",
			},
		},
	}
	for _, testcase := range testCases {
		s.T().Run(testcase.Name, func(t *testing.T) {
			r := httptest.NewRequest(testcase.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.control.GetAllUsers(ctx)
			s.NoError(err)

			s.Equal(testcase.ExpectedStatusCode, w.Result().StatusCode)
		})
	}
}

func (s *suiteUsers) TestGetUser_NotFound() {
	userID := "nonexistent"
	s.mock.On("FindByID", userID).Return(nil, errors.New("User not found"))

	r := httptest.NewRequest("GET", "/user/"+userID, nil)
	w := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(r, w)
	ctx.SetParamNames("id")
	ctx.SetParamValues(userID)

	err := s.control.GetAllUsers(ctx)
	s.Error(err) // Periksa bahwa kesalahan terjadi
	s.Equal(http.StatusNotFound, w.Result().StatusCode)
}

func (s *suiteUsers) TestCreateUser() {
	mockUser := model.User{
		Email:    "",
		Password: "",
	}
	s.mock.On("Create", mockUser).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       model.User
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			true,
			model.User{
				Email:    "admin@gmail.com",
				Password: "admin123",
			},
		},
	}
	for _, testcase := range testCases {
		s.T().Run(testcase.Name, func(t *testing.T) {
			r := httptest.NewRequest(testcase.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.control.CreateUser(ctx)
			s.NoError(err)

			s.Equal(testcase.ExpectedStatusCode, w.Result().StatusCode)
		})
	}
}

func (s *suiteUsers) TestCreateUser_InvalidInput() {
	invalidUser := model.User{
		Email:    "",
		Password: "password123",
	}
	s.mock.On("Create", invalidUser).Return(errors.New("Invalid input"))

	r := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(r, w)

	err := s.control.CreateUser(ctx)
	s.Error(err) // Periksa bahwa kesalahan terjadi
	s.Equal(http.StatusOK, w.Result().StatusCode)
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
