// Code generated by mockery v2.2.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// StatusClient is an autogenerated mock type for the StatusClient type
type StatusClient struct {
	mock.Mock
}

// Status provides a mock function with given fields:
func (_m *StatusClient) Status() client.StatusWriter {
	ret := _m.Called()

	var r0 client.StatusWriter
	if rf, ok := ret.Get(0).(func() client.StatusWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.StatusWriter)
		}
	}

	return r0
}