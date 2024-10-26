// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	entity "mlm/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindNodeByReferral provides a mock function with given fields: referral
func (_m *Repository) FindNodeByReferral(referral string) (entity.Node, error) {
	ret := _m.Called(referral)

	if len(ret) == 0 {
		panic("no return value specified for FindNodeByReferral")
	}

	var r0 entity.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Node, error)); ok {
		return rf(referral)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Node); ok {
		r0 = rf(referral)
	} else {
		r0 = ret.Get(0).(entity.Node)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(referral)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
