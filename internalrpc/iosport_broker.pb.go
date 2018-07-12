// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internalrpc/iosport_broker.proto

package internalrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bet "github.com/bitgaming/go-protobuf-schema/sportsbook/bet"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IosportBroker service

type IosportBrokerClient interface {
	PlaceBet(ctx context.Context, in *bet.BetRequest, opts ...grpc.CallOption) (*bet.BetResponse, error)
	CancelBet(ctx context.Context, in *bet.BetRequest, opts ...grpc.CallOption) (*bet.BetResponse, error)
}

type iosportBrokerClient struct {
	cc *grpc.ClientConn
}

func NewIosportBrokerClient(cc *grpc.ClientConn) IosportBrokerClient {
	return &iosportBrokerClient{cc}
}

func (c *iosportBrokerClient) PlaceBet(ctx context.Context, in *bet.BetRequest, opts ...grpc.CallOption) (*bet.BetResponse, error) {
	out := new(bet.BetResponse)
	err := grpc.Invoke(ctx, "/internalrpc.IosportBroker/PlaceBet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iosportBrokerClient) CancelBet(ctx context.Context, in *bet.BetRequest, opts ...grpc.CallOption) (*bet.BetResponse, error) {
	out := new(bet.BetResponse)
	err := grpc.Invoke(ctx, "/internalrpc.IosportBroker/CancelBet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IosportBroker service

type IosportBrokerServer interface {
	PlaceBet(context.Context, *bet.BetRequest) (*bet.BetResponse, error)
	CancelBet(context.Context, *bet.BetRequest) (*bet.BetResponse, error)
}

func RegisterIosportBrokerServer(s *grpc.Server, srv IosportBrokerServer) {
	s.RegisterService(&_IosportBroker_serviceDesc, srv)
}

func _IosportBroker_PlaceBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bet.BetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IosportBrokerServer).PlaceBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalrpc.IosportBroker/PlaceBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IosportBrokerServer).PlaceBet(ctx, req.(*bet.BetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IosportBroker_CancelBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bet.BetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IosportBrokerServer).CancelBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalrpc.IosportBroker/CancelBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IosportBrokerServer).CancelBet(ctx, req.(*bet.BetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IosportBroker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internalrpc.IosportBroker",
	HandlerType: (*IosportBrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBet",
			Handler:    _IosportBroker_PlaceBet_Handler,
		},
		{
			MethodName: "CancelBet",
			Handler:    _IosportBroker_CancelBet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalrpc/iosport_broker.proto",
}

func init() { proto.RegisterFile("internalrpc/iosport_broker.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xce, 0x41, 0xaa, 0xc2, 0x30,
	0x10, 0xc6, 0xf1, 0xb7, 0x7a, 0x68, 0x44, 0x94, 0xac, 0xa4, 0x22, 0x88, 0x07, 0x48, 0x45, 0x6f,
	0x10, 0x57, 0xee, 0xc4, 0x0b, 0x48, 0x26, 0x0c, 0x52, 0x9a, 0x66, 0xe2, 0xcc, 0x78, 0x7f, 0x69,
	0xbb, 0xb0, 0x4b, 0x17, 0xb3, 0x19, 0x7e, 0x7f, 0xf8, 0xcc, 0xbe, 0xc9, 0x8a, 0x9c, 0x43, 0xe2,
	0x12, 0xeb, 0x86, 0xa4, 0x10, 0xeb, 0x03, 0x98, 0x5a, 0x64, 0x57, 0x98, 0x94, 0xec, 0x62, 0x22,
	0xaa, 0xcd, 0x00, 0x04, 0x88, 0xda, 0x1a, 0x50, 0xfb, 0x1b, 0xd9, 0x89, 0xcd, 0xf2, 0x3a, 0xe6,
	0x7e, 0xa8, 0x6d, 0x6d, 0x66, 0xb7, 0x14, 0x22, 0x7a, 0x54, 0xbb, 0x72, 0x3d, 0xf4, 0xa8, 0x77,
	0x7c, 0xbd, 0x51, 0xb4, 0x5a, 0x7f, 0x1f, 0x52, 0x28, 0x0b, 0x1e, 0xfe, 0xec, 0xd1, 0xcc, 0x2f,
	0x21, 0x47, 0x4c, 0xbf, 0x16, 0x7e, 0x67, 0xb6, 0x91, 0x3a, 0x07, 0x8d, 0x3e, 0x43, 0x87, 0x29,
	0x80, 0xb8, 0xc9, 0x58, 0xf8, 0x1f, 0x96, 0x9d, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xc5,
	0xa4, 0xf7, 0xe4, 0x00, 0x00, 0x00,
}
