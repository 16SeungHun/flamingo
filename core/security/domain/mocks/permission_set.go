// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// PermissionSet is an autogenerated mock type for the PermissionSet type
type PermissionSet struct {
	mock.Mock
}

// Permissions provides a mock function with given fields:
func (_m *PermissionSet) Permissions() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// NewPermissionSet creates a new instance of PermissionSet. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPermissionSet(t testing.TB) *PermissionSet {
	mock := &PermissionSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
