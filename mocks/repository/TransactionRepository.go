// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "assignment-golang-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// GetWithParams provides a mock function with given fields: sortBy, sortDirection, searchQuery, limit, walletId
func (_m *TransactionRepository) GetWithParams(sortBy string, sortDirection string, searchQuery string, limit int, walletId int) ([]*entity.Transaction, error) {
	ret := _m.Called(sortBy, sortDirection, searchQuery, limit, walletId)

	var r0 []*entity.Transaction
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) []*entity.Transaction); ok {
		r0 = rf(sortBy, sortDirection, searchQuery, limit, walletId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int, int) error); ok {
		r1 = rf(sortBy, sortDirection, searchQuery, limit, walletId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
