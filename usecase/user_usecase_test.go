package usecase_test

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	mockUtils "assignment-golang-backend/mocks/hashutils"
	mocks "assignment-golang-backend/mocks/repository"
	"assignment-golang-backend/sentinelerrors"
	"assignment-golang-backend/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	testCases := []struct {
		name               string
		requestDto         *dto.UserDto
		requestUser        *entity.User
		mockIsUserExist    bool
		mockHashedPassword string
		mockHashedError    error
		mockResult         *entity.User
		mockErrResult      error
		expectedError      error
		expectedResult     *entity.User
	}{
		{
			name:               "Should return nil error and valid result when register request is valid and successful",
			requestDto:         &dto.UserDto{},
			requestUser:        &entity.User{},
			mockIsUserExist:    false,
			mockHashedPassword: "",
			mockHashedError:    nil,
			mockResult:         &entity.User{},
			mockErrResult:      nil,
			expectedError:      nil,
			expectedResult:     &entity.User{},
		},
		{
			name:               "Should return error and nil result when register request email already exists",
			requestDto:         &dto.UserDto{},
			requestUser:        &entity.User{},
			mockIsUserExist:    true,
			mockHashedPassword: "",
			mockHashedError:    nil,
			mockResult:         &entity.User{},
			mockErrResult:      nil,
			expectedError:      sentinelerrors.ErrEmailAlreadyExists,
			expectedResult:     nil,
		},
		{
			name:               "Should return error and nil result when hashUtils fails to hash",
			requestDto:         &dto.UserDto{},
			requestUser:        &entity.User{},
			mockIsUserExist:    false,
			mockHashedPassword: "",
			mockHashedError:    errors.New("error"),
			mockResult:         &entity.User{},
			mockErrResult:      nil,
			expectedError:      errors.New("error"),
			expectedResult:     nil,
		},
	}

	for _, testCase := range testCases {
		mockRepo := mocks.NewUserRepository(t)
		mockHashUtils := mockUtils.NewHashUtils(t)
		useCase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepo,
			HashUtil:       mockHashUtils,
		})
		mockRepo.On("IsUserWithEmailExist", testCase.requestDto.Email).Return(testCase.mockIsUserExist)

		if !(testCase.mockIsUserExist) {
			mockHashUtils.On("HashAndSalt", testCase.requestDto.Password).Return(testCase.mockHashedPassword, testCase.mockHashedError)
		}

		if testCase.mockHashedError == nil && !(testCase.mockIsUserExist) {
			mockRepo.On("Register", testCase.requestUser).Return(testCase.mockResult, testCase.mockErrResult)
		}

		res, err := useCase.Register(testCase.requestDto)

		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedResult, res)
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		name               string
		requestDto         *dto.UserDto
		mockGetUserByEmail *entity.User
		mockGetuserErr     error
		mockCompareResult  bool
		mockTokenResult    *dto.LoginResponse
		mockTokenError     error
		expectedError      error
		expectedResult     *dto.LoginResponse
	}{
		{
			name:               "Should Return nil error and login response when user exists and compare succesful",
			requestDto:         &dto.UserDto{},
			mockGetUserByEmail: &entity.User{},
			mockGetuserErr:     nil,
			mockCompareResult:  true,
			mockTokenResult:    &dto.LoginResponse{},
			mockTokenError:     nil,
			expectedError:      nil,
			expectedResult:     &dto.LoginResponse{},
		},
		{
			name:               "Should Return error and nil login response when does not exists",
			requestDto:         &dto.UserDto{},
			mockGetUserByEmail: nil,
			mockGetuserErr:     sentinelerrors.ErrEmailNotExists,
			mockCompareResult:  true,
			mockTokenResult:    &dto.LoginResponse{},
			mockTokenError:     nil,
			expectedError:      sentinelerrors.ErrEmailNotExists,
			expectedResult:     nil,
		},
		{
			name:               "Should Return error and nil login response when password compare does not match",
			requestDto:         &dto.UserDto{},
			mockGetUserByEmail: &entity.User{},
			mockGetuserErr:     nil,
			mockCompareResult:  false,
			mockTokenResult:    &dto.LoginResponse{},
			mockTokenError:     nil,
			expectedError:      sentinelerrors.ErrInvalidPassword,
			expectedResult:     nil,
		},
		{
			name:               "Should Return error and nil login response when password compare does not match",
			requestDto:         &dto.UserDto{},
			mockGetUserByEmail: &entity.User{},
			mockGetuserErr:     nil,
			mockCompareResult:  false,
			mockTokenResult:    nil,
			mockTokenError:     sentinelerrors.ErrInvalidPassword,
			expectedError:      sentinelerrors.ErrInvalidPassword,
			expectedResult:     nil,
		},
		{
			name:               "Should Return error and nil login response when fail to generate access token",
			requestDto:         &dto.UserDto{},
			mockGetUserByEmail: &entity.User{},
			mockGetuserErr:     nil,
			mockCompareResult:  true,
			mockTokenResult:    nil,
			mockTokenError:     errors.New("dummy error"),
			expectedError:      errors.New("dummy error"),
			expectedResult:     nil,
		},
	}

	for _, testCase := range testCases {
		mockRepo := mocks.NewUserRepository(t)
		mockHashUtils := mockUtils.NewHashUtils(t)
		useCase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepo,
			HashUtil:       mockHashUtils,
		})

		mockRepo.On("GetUserByEmail", testCase.requestDto.Email).Return(testCase.mockGetUserByEmail, testCase.mockGetuserErr)

		if testCase.mockGetuserErr == nil {
			mockHashUtils.On("ComparePassword", testCase.mockGetUserByEmail.Password, testCase.requestDto.Password).Return(testCase.mockCompareResult)
		}

		if testCase.mockGetuserErr == nil && testCase.mockCompareResult {
			mockHashUtils.On("GenerateAccessToken", testCase.mockGetUserByEmail).Return(testCase.mockTokenResult, testCase.mockTokenError)
		}

		res, err := useCase.Login(testCase.requestDto)

		assert.Equal(t, testCase.expectedResult, res)
		assert.Equal(t, testCase.expectedError, err)
	}
}
