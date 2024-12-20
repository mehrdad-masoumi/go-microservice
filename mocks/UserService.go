// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	dto "mlm/dto"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: request
func (_m *UserService) Create(request dto.UserCreateRequest) (dto.UserCreateResponse, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 dto.UserCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.UserCreateRequest) (dto.UserCreateResponse, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(dto.UserCreateRequest) dto.UserCreateResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(dto.UserCreateResponse)
	}

	if rf, ok := ret.Get(1).(func(dto.UserCreateRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
