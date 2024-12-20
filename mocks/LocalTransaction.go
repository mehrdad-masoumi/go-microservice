// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LocalTransaction is an autogenerated mock type for the LocalTransaction type
type LocalTransaction struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *LocalTransaction) Execute() error {
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

// NewLocalTransaction creates a new instance of LocalTransaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLocalTransaction(t interface {
	mock.TestingT
	Cleanup(func())
}) *LocalTransaction {
	mock := &LocalTransaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
