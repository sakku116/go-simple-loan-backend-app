// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	context "context"
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"
)

// IFileRepo is an autogenerated mock type for the IFileRepo type
type IFileRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, filename, bucketname
func (_m *IFileRepo) Delete(ctx context.Context, filename string, bucketname string) error {
	ret := _m.Called(ctx, filename, bucketname)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, filename, bucketname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUrl provides a mock function with given fields: ctx, filename, bucketname, download
func (_m *IFileRepo) GetUrl(ctx context.Context, filename string, bucketname string, download bool) (string, error) {
	ret := _m.Called(ctx, filename, bucketname, download)

	if len(ret) == 0 {
		panic("no return value specified for GetUrl")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) (string, error)); ok {
		return rf(ctx, filename, bucketname, download)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) string); ok {
		r0 = rf(ctx, filename, bucketname, download)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool) error); ok {
		r1 = rf(ctx, filename, bucketname, download)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: ctx, file, filename, bucketname
func (_m *IFileRepo) Upload(ctx context.Context, file *multipart.FileHeader, filename string, bucketname string) (string, error) {
	ret := _m.Called(ctx, file, filename, bucketname)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *multipart.FileHeader, string, string) (string, error)); ok {
		return rf(ctx, file, filename, bucketname)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *multipart.FileHeader, string, string) string); ok {
		r0 = rf(ctx, file, filename, bucketname)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *multipart.FileHeader, string, string) error); ok {
		r1 = rf(ctx, file, filename, bucketname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIFileRepo creates a new instance of IFileRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIFileRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *IFileRepo {
	mock := &IFileRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}