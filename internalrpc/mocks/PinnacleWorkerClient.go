// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "golang.org/x/net/context"
import empty "github.com/golang/protobuf/ptypes/empty"
import grpc "google.golang.org/grpc"
import internalrpc "github.com/bitgaming/go-protobuf-schema/internalrpc"
import mock "github.com/stretchr/testify/mock"

// PinnacleWorkerClient is an autogenerated mock type for the PinnacleWorkerClient type
type PinnacleWorkerClient struct {
	mock.Mock
}

// BlockSpecialMarket provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) BlockSpecialMarket(ctx context.Context, in *internalrpc.SpecialMarketRequest, opts ...grpc.CallOption) (*internalrpc.BlockedSpecialMarketsList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *internalrpc.BlockedSpecialMarketsList
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.SpecialMarketRequest, ...grpc.CallOption) *internalrpc.BlockedSpecialMarketsList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.BlockedSpecialMarketsList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.SpecialMarketRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockSpecialMarketsForEvent provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) BlockSpecialMarketsForEvent(ctx context.Context, in *internalrpc.SpecialMarketsForEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.SpecialMarketsForEventRequest, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.SpecialMarketsForEventRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisablePinnacleLive provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) DisablePinnacleLive(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *empty.Empty, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
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

// EnablePinnacleLive provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) EnablePinnacleLive(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *empty.Empty, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
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

// GetAllBlockedSpecialMarkets provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) GetAllBlockedSpecialMarkets(ctx context.Context, in *internalrpc.SpecialMarketsForEventRequest, opts ...grpc.CallOption) (*internalrpc.BlockedSpecialMarketsList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *internalrpc.BlockedSpecialMarketsList
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.SpecialMarketsForEventRequest, ...grpc.CallOption) *internalrpc.BlockedSpecialMarketsList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.BlockedSpecialMarketsList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.SpecialMarketsForEventRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PinnacleLiveStatus provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) PinnacleLiveStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*internalrpc.PinnacleLiveStatusReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *internalrpc.PinnacleLiveStatusReply
	if rf, ok := ret.Get(0).(func(context.Context, *empty.Empty, ...grpc.CallOption) *internalrpc.PinnacleLiveStatusReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.PinnacleLiveStatusReply)
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

// UnblockSpecialMarket provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) UnblockSpecialMarket(ctx context.Context, in *internalrpc.SpecialMarketRequest, opts ...grpc.CallOption) (*internalrpc.BlockedSpecialMarketsList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *internalrpc.BlockedSpecialMarketsList
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.SpecialMarketRequest, ...grpc.CallOption) *internalrpc.BlockedSpecialMarketsList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalrpc.BlockedSpecialMarketsList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.SpecialMarketRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnblockSpecialMarketsForEvent provides a mock function with given fields: ctx, in, opts
func (_m *PinnacleWorkerClient) UnblockSpecialMarketsForEvent(ctx context.Context, in *internalrpc.SpecialMarketsForEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *internalrpc.SpecialMarketsForEventRequest, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalrpc.SpecialMarketsForEventRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
