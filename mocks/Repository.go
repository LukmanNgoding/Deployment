// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "main.go/features/user/domain"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ID
func (_m *Repository) Delete(ID uint) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ID
func (_m *Repository) Get(ID uint) (domain.Core, error) {
	ret := _m.Called(ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(uint) domain.Core); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() ([]domain.Core, error) {
	ret := _m.Called()

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func() []domain.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newUser
func (_m *Repository) Insert(newUser domain.Core) (domain.Core, error) {
	ret := _m.Called(newUser)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: newUser
func (_m *Repository) Login(newUser domain.Core) (domain.Core, error) {
	ret := _m.Called(newUser)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: updateData
func (_m *Repository) Update(updateData domain.Core) (domain.Core, error) {
	ret := _m.Called(updateData)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(updateData)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
