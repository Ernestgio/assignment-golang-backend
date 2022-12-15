// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "assignment-golang-backend/dto"
	entity "assignment-golang-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// HashUtils is an autogenerated mock type for the HashUtils type
type HashUtils struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: hashedPwd, inputPwd
func (_m *HashUtils) ComparePassword(hashedPwd string, inputPwd string) bool {
	ret := _m.Called(hashedPwd, inputPwd)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hashedPwd, inputPwd)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GenerateAccessToken provides a mock function with given fields: user
func (_m *HashUtils) GenerateAccessToken(user *entity.User) (*dto.LoginResponse, error) {
	ret := _m.Called(user)

	var r0 *dto.LoginResponse
	if rf, ok := ret.Get(0).(func(*entity.User) *dto.LoginResponse); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashAndSalt provides a mock function with given fields: password
func (_m *HashUtils) HashAndSalt(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: tokenString
func (_m *HashUtils) ValidateToken(tokenString string) (*dto.UserClaim, error) {
	ret := _m.Called(tokenString)

	var r0 *dto.UserClaim
	if rf, ok := ret.Get(0).(func(string) *dto.UserClaim); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHashUtils interface {
	mock.TestingT
	Cleanup(func())
}

// NewHashUtils creates a new instance of HashUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHashUtils(t mockConstructorTestingTNewHashUtils) *HashUtils {
	mock := &HashUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}