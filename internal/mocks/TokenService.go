// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields: user
func (_m *TokenService) CreateToken(user *domain.User) (string, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User) (string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: token
func (_m *TokenService) VerifyToken(token string) (*domain.TokenPayload, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 *domain.TokenPayload
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.TokenPayload, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.TokenPayload); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.TokenPayload)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenService creates a new instance of TokenService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenService {
	mock := &TokenService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
