// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "golang.org/x/net/context"
import empty "github.com/golang/protobuf/ptypes/empty"
import grpc "google.golang.org/grpc"
import internalrpc "github.com/bitgaming/go-protobuf-schema/internalrpc"
import mock "github.com/stretchr/testify/mock"

// SportsbookAPIClient is an autogenerated mock type for the SportsbookAPIClient type
type SportsbookAPIClient struct {
	mock.Mock
}

// GetFullLockingStatus provides a mock function with given fields: ctx, in, opts
func (_m *SportsbookAPIClient) GetFullLockingStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*internalrpc.FullLockingStatusReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *internalrpc.FullLockingStatusReply
	if rf, ok := ret.Get(0).(func(context.Context, *empty.Empty, ...grpc.CallOption) *internalrpc.FullLockingStatusReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.FullLockingStatusReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *empty.Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOverrideSelectionStatus provides a mock function with given fields: ctx, in, opts
func (_m *SportsbookAPIClient) GetOverrideSelectionStatus(ctx context.Context, in *internalrpc.OverrideSelectionStatusRequest, opts ...grpc.CallOption) (*internalrpc.OverrideSelectionStatusReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *internalrpc.OverrideSelectionStatusReply
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.OverrideSelectionStatusRequest, ...grpc.CallOption) *internalrpc.OverrideSelectionStatusReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.OverrideSelectionStatusReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.OverrideSelectionStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayerInfo provides a mock function with given fields: ctx, in, opts
func (_m *SportsbookAPIClient) GetPlayerInfo(ctx context.Context, in *internalrpc.PlayerRequest, opts ...grpc.CallOption) (*internalrpc.PlayerReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *internalrpc.PlayerReply
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.PlayerRequest, ...grpc.CallOption) *internalrpc.PlayerReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.PlayerReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.PlayerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLockingStatus provides a mock function with given fields: ctx, in, opts
func (_m *SportsbookAPIClient) SetLockingStatus(ctx context.Context, in *internalrpc.SetLockingStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.SetLockingStatusRequest, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.SetLockingStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}