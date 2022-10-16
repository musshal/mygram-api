// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "mygram-api/domain"

	mock "github.com/stretchr/testify/mock"
)

// PhotoUseCase is an autogenerated mock type for the PhotoUseCase type
type PhotoUseCase struct {
	mock.Mock
}

// AddPhoto provides a mock function with given fields: _a0, _a1
func (_m *PhotoUseCase) AddPhoto(_a0 context.Context, _a1 *domain.User) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePhoto provides a mock function with given fields: ctx, id
func (_m *PhotoUseCase) DeletePhoto(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPhotos provides a mock function with given fields: ctx
func (_m *PhotoUseCase) GetPhotos(ctx context.Context) ([]domain.Photo, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Photo
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Photo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Photo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePhoto provides a mock function with given fields: ctx, photo
func (_m *PhotoUseCase) UpdatePhoto(ctx context.Context, photo *domain.Photo) error {
	ret := _m.Called(ctx, photo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Photo) error); ok {
		r0 = rf(ctx, photo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPhotoUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewPhotoUseCase creates a new instance of PhotoUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPhotoUseCase(t mockConstructorTestingTNewPhotoUseCase) *PhotoUseCase {
	mock := &PhotoUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
