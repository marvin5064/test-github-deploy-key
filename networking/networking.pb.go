// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networking/networking.proto

/*
Package networking is a generated protocol buffer package.

It is generated from these files:
	networking/networking.proto
	networking/pusher.proto

It has these top-level messages:
	Frame
	Pusher
	Metadata
*/
package networking

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Frame shall be used for pusher messages
type Frame struct {
	// client shall allocate buffer of length bytes. Even if sequence is out
	// of order, data can always be inserted with offset.
	Uuid     *string   `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Offset   *uint32   `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Length   *uint32   `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
	Data     *string   `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	Metadata *Metadata `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	// going to deprecated
	MessageType      *string `protobuf:"bytes,5,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Frame) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Frame) GetOffset() uint32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *Frame) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *Frame) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *Frame) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Frame) GetMessageType() string {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return ""
}

func init() {
	proto.RegisterType((*Frame)(nil), "networking.Frame")
}

func init() { proto.RegisterFile("networking/networking.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x4e, 0x86, 0x30,
	0x14, 0x85, 0x53, 0x05, 0xa2, 0x17, 0x5d, 0x1a, 0xa3, 0x0d, 0x3a, 0xa0, 0x13, 0x13, 0x1a, 0x1f,
	0xc1, 0xc1, 0xcd, 0x85, 0xb8, 0x9b, 0x22, 0x97, 0x42, 0xa4, 0xb4, 0x69, 0x2f, 0x31, 0xbc, 0x97,
	0x0f, 0x68, 0xa0, 0xe4, 0x87, 0xed, 0x9c, 0xef, 0x9e, 0x7c, 0xb9, 0x70, 0x3f, 0x22, 0xfd, 0x1a,
	0xf7, 0xd3, 0x8f, 0xea, 0x79, 0x8f, 0xa5, 0x75, 0x86, 0x0c, 0x87, 0x9d, 0x64, 0x77, 0x87, 0xa1,
	0x9d, 0x7c, 0x87, 0x2e, 0x8c, 0x9e, 0xfe, 0x18, 0xc4, 0xef, 0x4e, 0x6a, 0xe4, 0x1c, 0xa2, 0x69,
	0xea, 0x1b, 0xc1, 0x72, 0x56, 0x5c, 0x56, 0x6b, 0xe6, 0xb7, 0x90, 0x98, 0xb6, 0xf5, 0x48, 0xe2,
	0x2c, 0x67, 0xc5, 0x75, 0xb5, 0xb5, 0x85, 0x0f, 0x38, 0x2a, 0xea, 0xc4, 0x79, 0xe0, 0xa1, 0x2d,
	0x8e, 0x46, 0x92, 0x14, 0x51, 0x70, 0x2c, 0x99, 0xbf, 0xc0, 0x85, 0x46, 0x92, 0x2b, 0x4f, 0x72,
	0x56, 0xa4, 0xaf, 0x37, 0xe5, 0xe1, 0xd7, 0x8f, 0xed, 0x56, 0x9d, 0x56, 0xfc, 0x11, 0xae, 0x34,
	0x7a, 0x2f, 0x15, 0x7e, 0xd1, 0x6c, 0x51, 0xc4, 0xab, 0x2d, 0xdd, 0xd8, 0xe7, 0x6c, 0xf1, 0xed,
	0x01, 0xb2, 0x6f, 0xa3, 0xcb, 0xba, 0x27, 0x25, 0x35, 0x0e, 0xb2, 0xf6, 0x07, 0xe7, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x90, 0x74, 0x88, 0x17, 0x01, 0x00, 0x00,
}
