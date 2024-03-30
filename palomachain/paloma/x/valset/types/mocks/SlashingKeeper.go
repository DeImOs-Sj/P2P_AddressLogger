// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// SlashingKeeper is an autogenerated mock type for the SlashingKeeper type
type SlashingKeeper struct {
	mock.Mock
}

// Jail provides a mock function with given fields: _a0, _a1
func (_m *SlashingKeeper) Jail(_a0 context.Context, _a1 types.ConsAddress) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Jail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ConsAddress) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// JailUntil provides a mock function with given fields: _a0, _a1, _a2
func (_m *SlashingKeeper) JailUntil(_a0 context.Context, _a1 types.ConsAddress, _a2 time.Time) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for JailUntil")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ConsAddress, time.Time) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSlashingKeeper creates a new instance of SlashingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSlashingKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *SlashingKeeper {
	mock := &SlashingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
