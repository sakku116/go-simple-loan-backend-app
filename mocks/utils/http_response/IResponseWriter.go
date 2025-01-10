// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// IResponseWriter is an autogenerated mock type for the IResponseWriter type
type IResponseWriter struct {
	mock.Mock
}

// HTTPCustomErr provides a mock function with given fields: ctx, err
func (_m *IResponseWriter) HTTPCustomErr(ctx *gin.Context, err error) {
	_m.Called(ctx, err)
}

// HTTPJson provides a mock function with given fields: ctx, data
func (_m *IResponseWriter) HTTPJson(ctx *gin.Context, data interface{}) {
	_m.Called(ctx, data)
}

// HTTPJsonErr provides a mock function with given fields: ctx, code, message, detail, data
func (_m *IResponseWriter) HTTPJsonErr(ctx *gin.Context, code int, message string, detail string, data interface{}) {
	_m.Called(ctx, code, message, detail, data)
}

// NewIResponseWriter creates a new instance of IResponseWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIResponseWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *IResponseWriter {
	mock := &IResponseWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
