// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"

	log "cosmossdk.io/log"

	mock "github.com/stretchr/testify/mock"

	types "github.com/palomachain/paloma/x/scheduler/types"
)

// Keeper is an autogenerated mock type for the Keeper type
type Keeper struct {
	mock.Mock
}

// AddNewJob provides a mock function with given fields: ctx, job
func (_m *Keeper) AddNewJob(ctx context.Context, job *types.Job) error {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for AddNewJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Job) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecuteJob provides a mock function with given fields: ctx, jobID, payload, senderAddress, contractAddr
func (_m *Keeper) ExecuteJob(ctx context.Context, jobID string, payload []byte, senderAddress cosmos_sdktypes.AccAddress, contractAddr cosmos_sdktypes.AccAddress) (uint64, error) {
	ret := _m.Called(ctx, jobID, payload, senderAddress, contractAddr)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteJob")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte, cosmos_sdktypes.AccAddress, cosmos_sdktypes.AccAddress) (uint64, error)); ok {
		return rf(ctx, jobID, payload, senderAddress, contractAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte, cosmos_sdktypes.AccAddress, cosmos_sdktypes.AccAddress) uint64); ok {
		r0 = rf(ctx, jobID, payload, senderAddress, contractAddr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte, cosmos_sdktypes.AccAddress, cosmos_sdktypes.AccAddress) error); ok {
		r1 = rf(ctx, jobID, payload, senderAddress, contractAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: ctx, addr
func (_m *Keeper) GetAccount(ctx context.Context, addr cosmos_sdktypes.AccAddress) authtypes.AccountI {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for GetAccount")
	}

	var r0 authtypes.AccountI
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.AccAddress) authtypes.AccountI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authtypes.AccountI)
		}
	}

	return r0
}

// GetJob provides a mock function with given fields: ctx, jobID
func (_m *Keeper) GetJob(ctx context.Context, jobID string) (*types.Job, error) {
	ret := _m.Called(ctx, jobID)

	if len(ret) == 0 {
		panic("no return value specified for GetJob")
	}

	var r0 *types.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*types.Job, error)); ok {
		return rf(ctx, jobID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Job); ok {
		r0 = rf(ctx, jobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logger provides a mock function with given fields: ctx
func (_m *Keeper) Logger(ctx context.Context) log.Logger {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Logger")
	}

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(context.Context) log.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// PreJobExecution provides a mock function with given fields: ctx, job
func (_m *Keeper) PreJobExecution(ctx context.Context, job *types.Job) error {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for PreJobExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Job) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScheduleNow provides a mock function with given fields: ctx, jobID, in, senderAddress, contractAddress
func (_m *Keeper) ScheduleNow(ctx context.Context, jobID string, in []byte, senderAddress cosmos_sdktypes.AccAddress, contractAddress cosmos_sdktypes.AccAddress) (uint64, error) {
	ret := _m.Called(ctx, jobID, in, senderAddress, contractAddress)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleNow")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte, cosmos_sdktypes.AccAddress, cosmos_sdktypes.AccAddress) (uint64, error)); ok {
		return rf(ctx, jobID, in, senderAddress, contractAddress)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte, cosmos_sdktypes.AccAddress, cosmos_sdktypes.AccAddress) uint64); ok {
		r0 = rf(ctx, jobID, in, senderAddress, contractAddress)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte, cosmos_sdktypes.AccAddress, cosmos_sdktypes.AccAddress) error); ok {
		r1 = rf(ctx, jobID, in, senderAddress, contractAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKeeper creates a new instance of Keeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *Keeper {
	mock := &Keeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}