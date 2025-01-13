// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	dto "backend/domain/dto"
	context "context"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// IUserUcase is an autogenerated mock type for the IUserUcase type
type IUserUcase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, payload
func (_m *IUserUcase) CreateUser(ctx context.Context, payload dto.CreateUserReq) (*dto.CreateUserRespData, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *dto.CreateUserRespData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.CreateUserReq) (*dto.CreateUserRespData, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.CreateUserReq) *dto.CreateUserRespData); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CreateUserRespData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.CreateUserReq) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, userUUID
func (_m *IUserUcase) DeleteUser(ctx context.Context, userUUID string) (*dto.DeleteUserRespData, error) {
	ret := _m.Called(ctx, userUUID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 *dto.DeleteUserRespData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*dto.DeleteUserRespData, error)); ok {
		return rf(ctx, userUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *dto.DeleteUserRespData); ok {
		r0 = rf(ctx, userUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.DeleteUserRespData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUUID provides a mock function with given fields: ctx, userUUID
func (_m *IUserUcase) GetByUUID(ctx context.Context, userUUID string) (*dto.GetUserByUUIDResp, error) {
	ret := _m.Called(ctx, userUUID)

	if len(ret) == 0 {
		panic("no return value specified for GetByUUID")
	}

	var r0 *dto.GetUserByUUIDResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*dto.GetUserByUUIDResp, error)); ok {
		return rf(ctx, userUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *dto.GetUserByUUIDResp); ok {
		r0 = rf(ctx, userUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetUserByUUIDResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserList provides a mock function with given fields: ctx, params
func (_m *IUserUcase) GetUserList(ctx context.Context, params dto.GetUserListReq) (*dto.GetUserListRespData, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetUserList")
	}

	var r0 *dto.GetUserListRespData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.GetUserListReq) (*dto.GetUserListRespData, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.GetUserListReq) *dto.GetUserListRespData); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetUserListRespData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.GetUserListReq) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCurrentLimit provides a mock function with given fields: ctx, userUUID, payload
func (_m *IUserUcase) UpdateCurrentLimit(ctx context.Context, userUUID string, payload dto.UpdateCurrentLimitReq) (*dto.UpdateCurrentLimitRespData, error) {
	ret := _m.Called(ctx, userUUID, payload)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCurrentLimit")
	}

	var r0 *dto.UpdateCurrentLimitRespData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, dto.UpdateCurrentLimitReq) (*dto.UpdateCurrentLimitRespData, error)); ok {
		return rf(ctx, userUUID, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, dto.UpdateCurrentLimitReq) *dto.UpdateCurrentLimitRespData); ok {
		r0 = rf(ctx, userUUID, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UpdateCurrentLimitRespData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, dto.UpdateCurrentLimitReq) error); ok {
		r1 = rf(ctx, userUUID, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, userUUID, payload
func (_m *IUserUcase) UpdateUser(ctx context.Context, userUUID string, payload dto.UpdateUserReq) (*dto.UpdateUserRespData, error) {
	ret := _m.Called(ctx, userUUID, payload)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *dto.UpdateUserRespData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, dto.UpdateUserReq) (*dto.UpdateUserRespData, error)); ok {
		return rf(ctx, userUUID, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, dto.UpdateUserReq) *dto.UpdateUserRespData); ok {
		r0 = rf(ctx, userUUID, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UpdateUserRespData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, dto.UpdateUserReq) error); ok {
		r1 = rf(ctx, userUUID, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadFacePhoto provides a mock function with given fields: ctx, userUUID, file
func (_m *IUserUcase) UploadFacePhoto(ctx context.Context, userUUID string, file *multipart.FileHeader) (*dto.UploadFacePhotoRespData, error) {
	ret := _m.Called(ctx, userUUID, file)

	if len(ret) == 0 {
		panic("no return value specified for UploadFacePhoto")
	}

	var r0 *dto.UploadFacePhotoRespData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *multipart.FileHeader) (*dto.UploadFacePhotoRespData, error)); ok {
		return rf(ctx, userUUID, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *multipart.FileHeader) *dto.UploadFacePhotoRespData); ok {
		r0 = rf(ctx, userUUID, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UploadFacePhotoRespData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *multipart.FileHeader) error); ok {
		r1 = rf(ctx, userUUID, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadKtpPhoto provides a mock function with given fields: ctx, userUUID, file
func (_m *IUserUcase) UploadKtpPhoto(ctx context.Context, userUUID string, file *multipart.FileHeader) (*dto.UploadKtpPhotoRespData, error) {
	ret := _m.Called(ctx, userUUID, file)

	if len(ret) == 0 {
		panic("no return value specified for UploadKtpPhoto")
	}

	var r0 *dto.UploadKtpPhotoRespData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *multipart.FileHeader) (*dto.UploadKtpPhotoRespData, error)); ok {
		return rf(ctx, userUUID, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *multipart.FileHeader) *dto.UploadKtpPhotoRespData); ok {
		r0 = rf(ctx, userUUID, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UploadKtpPhotoRespData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *multipart.FileHeader) error); ok {
		r1 = rf(ctx, userUUID, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUserUcase creates a new instance of IUserUcase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserUcase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserUcase {
	mock := &IUserUcase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}