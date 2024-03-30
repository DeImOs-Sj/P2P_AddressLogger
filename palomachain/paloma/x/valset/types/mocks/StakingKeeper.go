// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// StakingKeeper is an autogenerated mock type for the StakingKeeper type
type StakingKeeper struct {
	mock.Mock
}

// IterateValidators provides a mock function with given fields: ctx, fn
func (_m *StakingKeeper) IterateValidators(ctx context.Context, fn func(int64, types.ValidatorI) bool) error {
	ret := _m.Called(ctx, fn)

	if len(ret) == 0 {
		panic("no return value specified for IterateValidators")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(int64, types.ValidatorI) bool) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Jail provides a mock function with given fields: ctx, consAddr
func (_m *StakingKeeper) Jail(ctx context.Context, consAddr cosmos_sdktypes.ConsAddress) error {
	ret := _m.Called(ctx, consAddr)

	if len(ret) == 0 {
		panic("no return value specified for Jail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ConsAddress) error); ok {
		r0 = rf(ctx, consAddr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validator provides a mock function with given fields: ctx, addr
func (_m *StakingKeeper) Validator(ctx context.Context, addr cosmos_sdktypes.ValAddress) (types.ValidatorI, error) {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for Validator")
	}

	var r0 types.ValidatorI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ValAddress) (types.ValidatorI, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ValAddress) types.ValidatorI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ValidatorI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmos_sdktypes.ValAddress) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStakingKeeper creates a new instance of StakingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStakingKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *StakingKeeper {
	mock := &StakingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}