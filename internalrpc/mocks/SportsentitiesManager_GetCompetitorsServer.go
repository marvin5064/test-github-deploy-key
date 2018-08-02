// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "golang.org/x/net/context"
import internalrpc "github.com/bitgaming/go-protobuf-schema/internalrpc"
import metadata "google.golang.org/grpc/metadata"
import mock "github.com/stretchr/testify/mock"

// SportsentitiesManager_GetCompetitorsServer is an autogenerated mock type for the SportsentitiesManager_GetCompetitorsServer type
type SportsentitiesManager_GetCompetitorsServer struct {
	mock.Mock
}

// Context provides a mock function with given fields:
func (_m *SportsentitiesManager_GetCompetitorsServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// RecvMsg provides a mock function with given fields: m
func (_m *SportsentitiesManager_GetCompetitorsServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *SportsentitiesManager_GetCompetitorsServer) Send(_a0 *internalrpc.GetCompetitorsReply) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internalrpc.GetCompetitorsReply) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendHeader provides a mock function with given fields: _a0
func (_m *SportsentitiesManager_GetCompetitorsServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *SportsentitiesManager_GetCompetitorsServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetHeader provides a mock function with given fields: _a0
func (_m *SportsentitiesManager_GetCompetitorsServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *SportsentitiesManager_GetCompetitorsServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}
