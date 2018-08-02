// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "golang.org/x/net/context"
import empty "github.com/golang/protobuf/ptypes/empty"
import internalrpc "github.com/bitgaming/go-protobuf-schema/internalrpc"
import mock "github.com/stretchr/testify/mock"

// SportsbookAPIServer is an autogenerated mock type for the SportsbookAPIServer type
type SportsbookAPIServer struct {
	mock.Mock
}

// GetFullLockingStatus provides a mock function with given fields: _a0, _a1
func (_m *SportsbookAPIServer) GetFullLockingStatus(_a0 context.Context, _a1 *empty.Empty) (*internalrpc.FullLockingStatusReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalrpc.FullLockingStatusReply
	if rf, ok := ret.Get(0).(func(context.Context, *empty.Empty) *internalrpc.FullLockingStatusReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.FullLockingStatusReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *empty.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOverrideSelectionStatus provides a mock function with given fields: _a0, _a1
func (_m *SportsbookAPIServer) GetOverrideSelectionStatus(_a0 context.Context, _a1 *internalrpc.OverrideSelectionStatusRequest) (*internalrpc.OverrideSelectionStatusReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalrpc.OverrideSelectionStatusReply
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.OverrideSelectionStatusRequest) *internalrpc.OverrideSelectionStatusReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.OverrideSelectionStatusReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.OverrideSelectionStatusRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayerInfo provides a mock function with given fields: _a0, _a1
func (_m *SportsbookAPIServer) GetPlayerInfo(_a0 context.Context, _a1 *internalrpc.PlayerRequest) (*internalrpc.PlayerReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalrpc.PlayerReply
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.PlayerRequest) *internalrpc.PlayerReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.PlayerReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.PlayerRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLockingStatus provides a mock function with given fields: _a0, _a1
func (_m *SportsbookAPIServer) SetLockingStatus(_a0 context.Context, _a1 *internalrpc.SetLockingStatusRequest) (*empty.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.SetLockingStatusRequest) *empty.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.SetLockingStatusRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
