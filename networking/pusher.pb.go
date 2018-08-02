// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networking/pusher.proto

package networking

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Pusher struct {
	Channel          *string   `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	Event            *string   `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
	Payload          []byte    `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	Metadata         *Metadata `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Pusher) Reset()                    { *m = Pusher{} }
func (m *Pusher) String() string            { return proto.CompactTextString(m) }
func (*Pusher) ProtoMessage()               {}
func (*Pusher) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Pusher) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *Pusher) GetEvent() string {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return ""
}

func (m *Pusher) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Pusher) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Metadata struct {
	Version          *uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	MessageType      *string `protobuf:"bytes,2,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
	ResourceKey      *uint32 `protobuf:"varint,3,opt,name=resource_key,json=resourceKey" json:"resource_key,omitempty"`
	PackageName      *string `protobuf:"bytes,4,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Metadata) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Metadata) GetMessageType() string {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return ""
}

func (m *Metadata) GetResourceKey() uint32 {
	if m != nil && m.ResourceKey != nil {
		return *m.ResourceKey
	}
	return 0
}

func (m *Metadata) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func init() {
	proto.RegisterType((*Pusher)(nil), "networking.Pusher")
	proto.RegisterType((*Metadata)(nil), "networking.Metadata")
}

func init() { proto.RegisterFile("networking/pusher.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x4e, 0xf3, 0x30,
	0x10, 0xc4, 0xe5, 0xef, 0xa3, 0x50, 0x36, 0xed, 0xc5, 0xaa, 0x84, 0x85, 0x38, 0x84, 0x9e, 0x72,
	0x0a, 0x88, 0x47, 0xe0, 0x8a, 0x40, 0xc8, 0xe2, 0x5e, 0x6d, 0xd3, 0x55, 0x1a, 0x25, 0xfe, 0x23,
	0xdb, 0x2d, 0xf2, 0x9d, 0x2b, 0xef, 0x8c, 0x9c, 0xc6, 0x94, 0xe3, 0xfc, 0x66, 0x34, 0x3b, 0x0b,
	0x37, 0x9a, 0xc2, 0xa7, 0x71, 0x7d, 0xa7, 0xdb, 0x07, 0x7b, 0xf0, 0x7b, 0x72, 0xb5, 0x75, 0x26,
	0x18, 0x0e, 0x67, 0x63, 0xfd, 0xc5, 0xe0, 0xf2, 0x7d, 0x34, 0xb9, 0x80, 0xab, 0x66, 0x8f, 0x5a,
	0xd3, 0x20, 0x58, 0xc9, 0xaa, 0x6b, 0x99, 0x25, 0x5f, 0xc1, 0x8c, 0x8e, 0xa4, 0x83, 0xf8, 0x37,
	0xf2, 0x93, 0x48, 0x79, 0x8b, 0x71, 0x30, 0xb8, 0x13, 0xff, 0x4b, 0x56, 0x2d, 0x64, 0x96, 0xfc,
	0x11, 0xe6, 0x8a, 0x02, 0xee, 0x30, 0xa0, 0x98, 0x95, 0xac, 0x2a, 0x9e, 0x56, 0xf5, 0xf9, 0x66,
	0xfd, 0x3a, 0x79, 0xf2, 0x37, 0xb5, 0xfe, 0x66, 0x30, 0xcf, 0x38, 0x15, 0x1f, 0xc9, 0xf9, 0xce,
	0xe8, 0x71, 0xc8, 0x52, 0x66, 0xc9, 0xef, 0x61, 0xa1, 0xc8, 0x7b, 0x6c, 0x69, 0x13, 0xa2, 0xa5,
	0x69, 0x4f, 0x31, 0xb1, 0x8f, 0x68, 0x29, 0x45, 0x1c, 0x79, 0x73, 0x70, 0x0d, 0x6d, 0x7a, 0x8a,
	0xe3, 0xb4, 0xa5, 0x2c, 0x32, 0x7b, 0xa1, 0x98, 0x22, 0x16, 0x9b, 0x3e, 0xb5, 0x68, 0x54, 0x24,
	0x2e, 0x4e, 0x2d, 0x13, 0x7b, 0x43, 0x45, 0xcf, 0x77, 0x70, 0xdb, 0x18, 0x55, 0x6f, 0xbb, 0xd0,
	0xa2, 0xa2, 0x01, 0xb7, 0xfe, 0xcf, 0x03, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x78, 0xb6,
	0xc7, 0x5a, 0x01, 0x00, 0x00,
}