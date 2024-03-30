// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	"context"
	consensus "github.com/palomachain/paloma/x/consensus/keeper/consensus"
	consensustypes "github.com/palomachain/paloma/x/consensus/types"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Queuer is an autogenerated mock type for the Queuer type
type Queuer struct {
	mock.Mock
}

// AddEvidence provides a mock function with given fields: ctx, id, evidence
func (_m *Queuer) AddEvidence(ctx context.Context, id uint64, evidence *consensustypes.Evidence) error {
	ret := _m.Called(ctx, id, evidence)

	if len(ret) == 0 {
		panic("no return value specified for AddEvidence")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *consensustypes.Evidence) error); ok {
		r0 = rf(ctx, id, evidence)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddSignature provides a mock function with given fields: ctx, id, signData
func (_m *Queuer) AddSignature(ctx context.Context, id uint64, signData *consensustypes.SignData) error {
	ret := _m.Called(ctx, id, signData)

	if len(ret) == 0 {
		panic("no return value specified for AddSignature")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *consensustypes.SignData) error); ok {
		r0 = rf(ctx, id, signData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChainInfo provides a mock function with given fields:
func (_m *Queuer) ChainInfo() (string, string) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ChainInfo")
	}

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func() (string, string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: _a0
func (_m *Queuer) GetAll(_a0 context.Context) ([]consensustypes.QueuedSignedMessageI, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []consensustypes.QueuedSignedMessageI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]consensustypes.QueuedSignedMessageI, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []consensustypes.QueuedSignedMessageI); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consensustypes.QueuedSignedMessageI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetErrorData provides a mock function with given fields: ctx, id
func (_m *Queuer) GetErrorData(ctx context.Context, id uint64) (*consensustypes.ErrorData, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetErrorData")
	}

	var r0 *consensustypes.ErrorData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*consensustypes.ErrorData, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *consensustypes.ErrorData); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consensustypes.ErrorData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMsgByID provides a mock function with given fields: ctx, id
func (_m *Queuer) GetMsgByID(ctx context.Context, id uint64) (consensustypes.QueuedSignedMessageI, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetMsgByID")
	}

	var r0 consensustypes.QueuedSignedMessageI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (consensustypes.QueuedSignedMessageI, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) consensustypes.QueuedSignedMessageI); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(consensustypes.QueuedSignedMessageI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicAccessData provides a mock function with given fields: ctx, id
func (_m *Queuer) GetPublicAccessData(ctx context.Context, id uint64) (*consensustypes.PublicAccessData, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPublicAccessData")
	}

	var r0 *consensustypes.PublicAccessData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*consensustypes.PublicAccessData, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *consensustypes.PublicAccessData); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consensustypes.PublicAccessData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: _a0, _a1, _a2
func (_m *Queuer) Put(_a0 context.Context, _a1 consensustypes.ConsensusMsg, _a2 *consensus.PutOptions) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, consensustypes.ConsensusMsg, *consensus.PutOptions) (uint64, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, consensustypes.ConsensusMsg, *consensus.PutOptions) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, consensustypes.ConsensusMsg, *consensus.PutOptions) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReassignValidator provides a mock function with given fields: ctx, id, val
func (_m *Queuer) ReassignValidator(ctx types.Context, id uint64, val string) error {
	ret := _m.Called(ctx, id, val)

	if len(ret) == 0 {
		panic("no return value specified for ReassignValidator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, string) error); ok {
		r0 = rf(ctx, id, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: _a0, _a1
func (_m *Queuer) Remove(_a0 context.Context, _a1 uint64) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetErrorData provides a mock function with given fields: ctx, id, data
func (_m *Queuer) SetErrorData(ctx context.Context, id uint64, data *consensustypes.ErrorData) error {
	ret := _m.Called(ctx, id, data)

	if len(ret) == 0 {
		panic("no return value specified for SetErrorData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *consensustypes.ErrorData) error); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPublicAccessData provides a mock function with given fields: ctx, id, data
func (_m *Queuer) SetPublicAccessData(ctx context.Context, id uint64, data *consensustypes.PublicAccessData) error {
	ret := _m.Called(ctx, id, data)

	if len(ret) == 0 {
		panic("no return value specified for SetPublicAccessData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *consensustypes.PublicAccessData) error); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQueuer creates a new instance of Queuer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueuer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Queuer {
	mock := &Queuer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
