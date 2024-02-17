// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	rest "k8s.io/client-go/rest"
)

// NewClientFunc is an autogenerated mock type for the NewClientFunc type
type NewClientFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: config, options
func (_m *NewClientFunc) Execute(config *rest.Config, options client.Options) (client.Client, error) {
	ret := _m.Called(config, options)

	var r0 client.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(*rest.Config, client.Options) (client.Client, error)); ok {
		return rf(config, options)
	}
	if rf, ok := ret.Get(0).(func(*rest.Config, client.Options) client.Client); ok {
		r0 = rf(config, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(*rest.Config, client.Options) error); ok {
		r1 = rf(config, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNewClientFunc creates a new instance of NewClientFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNewClientFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *NewClientFunc {
	mock := &NewClientFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}