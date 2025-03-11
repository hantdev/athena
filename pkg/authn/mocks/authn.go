package mocks

import (
	context "context"

	authn "github.com/hantdev/athena/pkg/authn"

	mock "github.com/stretchr/testify/mock"
)

// Authentication is an autogenerated mock type for the Authentication type
type Authentication struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: ctx, token
func (_m *Authentication) Authenticate(ctx context.Context, token string) (authn.Session, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for Authenticate")
	}

	var r0 authn.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (authn.Session, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) authn.Session); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(authn.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthentication creates a new instance of Authentication. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthentication(t interface {
	mock.TestingT
	Cleanup(func())
}) *Authentication {
	mock := &Authentication{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
