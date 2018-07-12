// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "golang.org/x/net/context"
import empty "github.com/golang/protobuf/ptypes/empty"
import internalrpc "github.com/bitgaming/go-protobuf-schema/internalrpc"
import mock "github.com/stretchr/testify/mock"

// RiskManagerServer is an autogenerated mock type for the RiskManagerServer type
type RiskManagerServer struct {
	mock.Mock
}

// GetAgentBettingOverlay provides a mock function with given fields: _a0, _a1
func (_m *RiskManagerServer) GetAgentBettingOverlay(_a0 context.Context, _a1 *internalrpc.GetBettingOverlayRequest) (*internalrpc.EntityBettingOverlays, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalrpc.EntityBettingOverlays
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.GetBettingOverlayRequest) *internalrpc.EntityBettingOverlays); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.EntityBettingOverlays)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.GetBettingOverlayRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDirectChildrenBettingOverlays provides a mock function with given fields: _a0, _a1
func (_m *RiskManagerServer) GetDirectChildrenBettingOverlays(_a0 context.Context, _a1 *internalrpc.GetBettingOverlayRequest) (*internalrpc.BettingOverlays, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalrpc.BettingOverlays
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.GetBettingOverlayRequest) *internalrpc.BettingOverlays); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.BettingOverlays)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.GetBettingOverlayRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayerBettingOverlay provides a mock function with given fields: _a0, _a1
func (_m *RiskManagerServer) GetPlayerBettingOverlay(_a0 context.Context, _a1 *internalrpc.GetBettingOverlayRequest) (*internalrpc.EntityBettingOverlays, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalrpc.EntityBettingOverlays
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.GetBettingOverlayRequest) *internalrpc.EntityBettingOverlays); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.EntityBettingOverlays)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.GetBettingOverlayRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveBettingOverlay provides a mock function with given fields: _a0, _a1
func (_m *RiskManagerServer) RemoveBettingOverlay(_a0 context.Context, _a1 *internalrpc.RemoveBettingOverlayRequest) (*empty.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.RemoveBettingOverlayRequest) *empty.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.RemoveBettingOverlayRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetBettingOverlay provides a mock function with given fields: _a0, _a1
func (_m *RiskManagerServer) SetBettingOverlay(_a0 context.Context, _a1 *internalrpc.BettingOverlay) (*internalrpc.SingleId, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalrpc.SingleId
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.BettingOverlay) *internalrpc.SingleId); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.SingleId)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.BettingOverlay) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}