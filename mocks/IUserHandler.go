// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// IUserHandler is an autogenerated mock type for the IUserHandler type
type IUserHandler struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: c
func (_m *IUserHandler) CreateUser(c *gin.Context) {
	_m.Called(c)
}

// DeleteUser provides a mock function with given fields: c
func (_m *IUserHandler) DeleteUser(c *gin.Context) {
	_m.Called(c)
}

// GetByUUID provides a mock function with given fields: c
func (_m *IUserHandler) GetByUUID(c *gin.Context) {
	_m.Called(c)
}

// GetMe provides a mock function with given fields: c
func (_m *IUserHandler) GetMe(c *gin.Context) {
	_m.Called(c)
}

// GetUserList provides a mock function with given fields: c
func (_m *IUserHandler) GetUserList(c *gin.Context) {
	_m.Called(c)
}

// UpdateCurrentLimit provides a mock function with given fields: c
func (_m *IUserHandler) UpdateCurrentLimit(c *gin.Context) {
	_m.Called(c)
}

// UpdateUser provides a mock function with given fields: c
func (_m *IUserHandler) UpdateUser(c *gin.Context) {
	_m.Called(c)
}

// UpdateUserMe provides a mock function with given fields: c
func (_m *IUserHandler) UpdateUserMe(c *gin.Context) {
	_m.Called(c)
}

// UploadFacePhoto provides a mock function with given fields: c
func (_m *IUserHandler) UploadFacePhoto(c *gin.Context) {
	_m.Called(c)
}

// UploadKtpPhoto provides a mock function with given fields: c
func (_m *IUserHandler) UploadKtpPhoto(c *gin.Context) {
	_m.Called(c)
}

// NewIUserHandler creates a new instance of IUserHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserHandler {
	mock := &IUserHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
