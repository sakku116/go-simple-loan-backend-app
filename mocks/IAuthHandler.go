// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// IAuthHandler is an autogenerated mock type for the IAuthHandler type
type IAuthHandler struct {
	mock.Mock
}

// CheckToken provides a mock function with given fields: ctx
func (_m *IAuthHandler) CheckToken(ctx *gin.Context) {
	_m.Called(ctx)
}

// Login provides a mock function with given fields: ctx
func (_m *IAuthHandler) Login(ctx *gin.Context) {
	_m.Called(ctx)
}

// LoginDev provides a mock function with given fields: ctx
func (_m *IAuthHandler) LoginDev(ctx *gin.Context) {
	_m.Called(ctx)
}

// RefreshToken provides a mock function with given fields: ctx
func (_m *IAuthHandler) RefreshToken(ctx *gin.Context) {
	_m.Called(ctx)
}

// Register provides a mock function with given fields: ctx
func (_m *IAuthHandler) Register(ctx *gin.Context) {
	_m.Called(ctx)
}

// NewIAuthHandler creates a new instance of IAuthHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthHandler {
	mock := &IAuthHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}