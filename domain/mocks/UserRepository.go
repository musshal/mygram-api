// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "mygram-api/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: ctx, user
func (_m *UserRepository) AddUser(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) CheckEmail(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckUser provides a mock function with given fields: ctx, id
func (_m *UserRepository) CheckUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckUsername provides a mock function with given fields: ctx, username
func (_m *UserRepository) CheckUsername(ctx context.Context, username string) error {
	ret := _m.Called(ctx, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUsers provides a mock function with given fields: ctx
func (_m *UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	ret := _m.Called(ctx)

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func(context.Context) []domain.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
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

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
