// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "golang.org/x/net/context"
import empty "github.com/golang/protobuf/ptypes/empty"
import internalrpc "github.com/bitgaming/go-protobuf-schema/internalrpc"
import metadata "google.golang.org/grpc/metadata"
import mock "github.com/stretchr/testify/mock"

// SportsentitiesManager_UpdateTranslationsClient is an autogenerated mock type for the SportsentitiesManager_UpdateTranslationsClient type
type SportsentitiesManager_UpdateTranslationsClient struct {
	mock.Mock
}

// CloseAndRecv provides a mock function with given fields:
func (_m *SportsentitiesManager_UpdateTranslationsClient) CloseAndRecv() (*empty.Empty, error) {
	ret := _m.Called()

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func() *empty.Empty); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseSend provides a mock function with given fields:
func (_m *SportsentitiesManager_UpdateTranslationsClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *SportsentitiesManager_UpdateTranslationsClient) Context() context.Context {
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

// Header provides a mock function with given fields:
func (_m *SportsentitiesManager_UpdateTranslationsClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecvMsg provides a mock function with given fields: m
func (_m *SportsentitiesManager_UpdateTranslationsClient) RecvMsg(m interface{}) error {
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
func (_m *SportsentitiesManager_UpdateTranslationsClient) Send(_a0 *internalrpc.TranslationMap) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internalrpc.TranslationMap) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *SportsentitiesManager_UpdateTranslationsClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trailer provides a mock function with given fields:
func (_m *SportsentitiesManager_UpdateTranslationsClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}