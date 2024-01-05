// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TraceExiter is an autogenerated mock type for the TraceExiter type
type TraceExiter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: err, additionalValues
func (_m *TraceExiter) Execute(err error, additionalValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err)
	_ca = append(_ca, additionalValues...)
	_m.Called(_ca...)
}

// NewTraceExiter creates a new instance of TraceExiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTraceExiter(t interface {
	mock.TestingT
	Cleanup(func())
}) *TraceExiter {
	mock := &TraceExiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
