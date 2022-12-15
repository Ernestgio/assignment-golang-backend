package handler_test

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/hashutils"
	mocks "assignment-golang-backend/mocks/usecase"
	"assignment-golang-backend/server"
	"assignment-golang-backend/testutils"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerRegister(t *testing.T) {
	t.Run("Should return 201 status code when request is valid", func(t *testing.T) {
		registerRequest := &dto.UserDto{
			Email:    "yo@email.com",
			Password: "123456",
		}
		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("Register", registerRequest).Return(&entity.User{}, nil)

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}

		req, _ := http.NewRequest("POST", "/register", testutils.MakeRequestBody(registerRequest))
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("Should return 500 status code when register request failed", func(t *testing.T) {
		registerRequest := &dto.UserDto{
			Email:    "yo@email.com",
			Password: "123456",
		}
		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("Register", registerRequest).Return(nil, errors.New("Failed to register"))

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}

		req, _ := http.NewRequest("POST", "/register", testutils.MakeRequestBody(registerRequest))
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("Should return 400 status code when register request body is not valid", func(t *testing.T) {
		registerRequest := &dto.UserDto{}
		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("Register", registerRequest).Return(&entity.User{}, nil)

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}

		req, _ := http.NewRequest("POST", "/register", testutils.MakeRequestBody(registerRequest))
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestHandlerLogin(t *testing.T) {
	t.Run("Should return 200 status code when request is valid", func(t *testing.T) {
		loginRequest := &dto.UserDto{
			Email:    "yo@email.com",
			Password: "123456",
		}

		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("Login", loginRequest).Return(&dto.LoginResponse{}, nil)

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}
		req, _ := http.NewRequest("POST", "/login", testutils.MakeRequestBody(loginRequest))
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should return 400 status code when request body is not valid", func(t *testing.T) {
		loginRequest := &dto.UserDto{}

		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("Login", loginRequest).Return(&dto.LoginResponse{}, nil)

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}
		req, _ := http.NewRequest("POST", "/login", testutils.MakeRequestBody(loginRequest))
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Should return 500 status code when usecase Failed", func(t *testing.T) {
		loginRequest := &dto.UserDto{
			Email:    "yo@email.com",
			Password: "123456",
		}

		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("Login", loginRequest).Return(nil, errors.New("Failed to login"))

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}
		req, _ := http.NewRequest("POST", "/login", testutils.MakeRequestBody(loginRequest))
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestHandlerGetUserById(t *testing.T) {
	t.Run("Should return 200 status code when request is valid", func(t *testing.T) {
		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("GetUserById", 1).Return(&entity.User{}, nil)

		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}
		req, _ := http.NewRequest("GET", "/users/", nil)
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should return 500 status code when usecase Failed", func(t *testing.T) {
		mockUserUsecase := new(mocks.UserUsecase)
		mockUserUsecase.On("GetUserById", 1).Return(nil, errors.New("Failed to get user"))

		cfg := &server.RouterConfig{
			UserUsecase: mockUserUsecase,
		}

		req, _ := http.NewRequest("GET", "/users/", nil)
		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
