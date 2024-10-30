// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CompensatingAction is an autogenerated mock type for the CompensatingAction type
type CompensatingAction struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *CompensatingAction) Execute() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCompensatingAction creates a new instance of CompensatingAction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCompensatingAction(t interface {
	mock.TestingT
	Cleanup(func())
}) *CompensatingAction {
	mock := &CompensatingAction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}