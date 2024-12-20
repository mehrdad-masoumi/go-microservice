// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	entity "mlm/entity"

	mock "github.com/stretchr/testify/mock"
)

// NodeRepository is an autogenerated mock type for the NodeRepository type
type NodeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: node
func (_m *NodeRepository) Create(node entity.Node) (entity.Node, error) {
	ret := _m.Called(node)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 entity.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Node) (entity.Node, error)); ok {
		return rf(node)
	}
	if rf, ok := ret.Get(0).(func(entity.Node) entity.Node); ok {
		r0 = rf(node)
	} else {
		r0 = ret.Get(0).(entity.Node)
	}

	if rf, ok := ret.Get(1).(func(entity.Node) error); ok {
		r1 = rf(node)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *NodeRepository) Delete(id uint) (bool, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNodeByReferral provides a mock function with given fields: referral
func (_m *NodeRepository) FindNodeByReferral(referral string) (entity.Node, error) {
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

// NewNodeRepository creates a new instance of NodeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNodeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *NodeRepository {
	mock := &NodeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
