// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ResponseItemType is an autogenerated mock type for the ResponseItemType type
type ResponseItemType struct {
	mock.Mock
}

type mockConstructorTestingTNewResponseItemType interface {
	mock.TestingT
	Cleanup(func())
}

// NewResponseItemType creates a new instance of ResponseItemType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResponseItemType(t mockConstructorTestingTNewResponseItemType) *ResponseItemType {
	mock := &ResponseItemType{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
